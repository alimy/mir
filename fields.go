package mir

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// TagName mir's struct tag string name
	TagName = "mir"
)

var (
	tagNotExist      tagError = "mir struct tag filed not exist"
	tagNoPathInfo    tagError = "mir struct tag not contains path info"
	tagNotValideType tagError = "not valide type, just struct and struct ptr is avalibale"
	tagMultGroupInfo tagError = "multiple group info in struct defined"
)

// tagError indicate error information
type tagError string

func (e tagError) Error() string {
	return string(e)
}

// tagBase indicate mir tag common information in struct tag string exclude handler
type tagBase struct {
	Method  string   // Method indicate method information in struct tag string
	Host    string   // Host indicate host information in struct tag string
	Path    string   // Path indicate path information in struct tag string
	Queries []string // Queries indicate path information in struct tag string
}

// tagInfo indicate mir tag information in struct tag string
type tagInfo struct {
	tagBase
	isGroup bool   // indicate whether a group field
	group   string // indicate group information in struct tag string
	handler string // indicate handler information in struct tag string
}

// TagField indicate mir tag info used to register route info to engine
type TagField struct {
	tagBase
	Handler interface{} // Handler indicate handler method
}

// TagMir contains TagFields by group
type TagMir struct {
	Group  string
	Fields []*TagField
}

// tagFieldsGroup indicate group-tagFields map
type tagFieldsGroup map[string]*TagMir

// TagMirFrom build TagMir items from entries slice
func TagMirFrom(entries ...interface{}) ([]*TagMir, error) {
	mergedTagMirs := make(tagFieldsGroup)
	for _, entry := range entries {
		if tagFields, err := tagMirFrom(entry); err == nil {
			// no actual field so just continue
			if len(tagFields.Fields) == 0 {
				continue
			}
			// merge tagFields by group
			if mergedFields, exist := mergedTagMirs[tagFields.Group]; exist {
				mergedFields.Fields = append(mergedFields.Fields, tagFields.Fields...)
			} else {
				mergedFields = &TagMir{
					Group:  tagFields.Group,
					Fields: make([]*TagField, 0, len(tagFields.Fields)),
				}
				mergedFields.Fields = append(mergedFields.Fields, tagFields.Fields...)
				mergedTagMirs[tagFields.Group] = mergedFields
			}
		} else {
			return nil, err
		}
	}

	// build result
	tagMirSlice := make([]*TagMir, 0, len(mergedTagMirs))
	for _, item := range mergedTagMirs {
		tagMirSlice = append(tagMirSlice, item)
	}
	return tagMirSlice, nil
}

// tagMirFrom build tagMir items from a entry
func tagMirFrom(entry interface{}) (*TagMir, error) {
	// used to find tagInfo
	entryType := reflect.TypeOf(entry)
	isPtr := false

	// get real entry type
	if entryType.Kind() == reflect.Ptr {
		isPtr = true
		entryType = entryType.Elem()
	}

	// entry must struct type
	if entryType.Kind() != reflect.Struct {
		return nil, tagNotValideType
	}

	// used to find method from T and lookup struct tag string of mir tag information
	var entryValue, entryPtrValue reflect.Value
	if isPtr {
		entryPtrValue = reflect.ValueOf(entry)
		entryValue = entryPtrValue.Elem()
	} else {
		entryValue = reflect.ValueOf(entry)
		entryPtrValue = entryValue.Addr()
	}

	// get tagMir from entryType and entryPtrType
	tagMir := &TagMir{Fields: make([]*TagField, 0)}
	for i := 0; i < entryType.NumField(); i++ {
		field := entryType.Field(i)
		switch tagInfo, err := tagInfoFrom(field); err {
		case nil:
			// group field so just parse group info.group info only have one field
			if tagInfo.isGroup {
				if tagMir.Group == "" {
					tagMir.Group = tagInfo.group
					break
				} else {
					return nil, tagMultGroupInfo
				}
			}
			// method field build tagField first
			if tagField, err := tagFieldFrom(entryValue, entryPtrValue, tagInfo); err == nil {
				tagMir.Fields = append(tagMir.Fields, tagField)
			} else {
				return nil, err
			}
		case tagNotExist:
			// normal field but had no mir tag info so just break to continue process next field
		default:
			return nil, err
		}
	}
	return tagMir, nil
}

// tagFieldFrom build tagField from entry and tagInfo
func tagFieldFrom(entryValue reflect.Value, entryPtrValue reflect.Value, tagInfo *tagInfo) (*TagField, error) {
	var m reflect.Value
	if v := entryValue.MethodByName(tagInfo.handler); v.IsValid() {
		m = v
	} else if v := entryPtrValue.MethodByName(tagInfo.handler); v.IsValid() {
		m = v
	} else {
		typeName := entryValue.Type().Name()
		err := fmt.Errorf("not found method %s in struct type %s or *%s", tagInfo.handler, typeName, typeName)
		return nil, err
	}
	tagField := &TagField{
		tagBase: tagInfo.tagBase,
		Handler: m.Interface()}
	return tagField, nil
}

// tagInfoFrom build tagInfo from field
func tagInfoFrom(field reflect.StructField) (*tagInfo, error) {
	var (
		method, host, path, handler string
		queries                     []string
	)

	// lookup mir tag info from struct field
	tag, exist := field.Tag.Lookup(TagName)
	if !exist {
		return nil, tagNotExist
	}

	// Skip leading space.
	i := 0
	for i < len(tag) && tag[i] == ' ' {
		i++
	}
	tag = tag[i:]

	// group info or method info
	switch field.Type.Name() {
	case "Group":
		return &tagInfo{group: tag, isGroup: true}, nil
	case "Get":
		method = MethodGet
	case "Put":
		method = MethodPut
	case "Post":
		method = MethodPost
	case "Delete":
		method = MethodDelete
	case "Head":
		method = MethodHead
	case "Options":
		method = MethodOptions
	case "Patch":
		method = MethodPatch
	case "Trace":
		method = MethodTrace
	case "Connect":
		method = MethodConnect
	case "Any":
		method = "ANY"
	}

	// host info
	if len(tag) > 2 && tag[0] == '/' && tag[1] == '/' {
		i := 2
		for i < len(tag) && tag[i] != '/' {
			i++
		}
		host = tag[2:i]
		tag = tag[i:]
	}

	// path info
	if len(tag) == 0 {
		return nil, tagNoPathInfo
	}
	i = 0
	for i < len(tag) && tag[i] != '?' && tag[i] != '#' {
		i++
	}
	path = tag[0:i]
	tag = tag[i:]

	// queries and handler info
	for len(tag) != 0 {
		switch tag[0] {
		case '#':
			i := 1
			for i < len(tag) && tag[i] != '?' {
				i++
			}
			handler = tag[1:i]
			tag = tag[i:]
		case '?':
			i := 1
			for i < len(tag) && tag[i] != '#' {
				i++
			}
			queryStr := tag[1:i]
			if queryStr != "" {
				queries = strings.Split(queryStr, "&")
			}
			tag = tag[i:]
		}
	}

	// check handler
	if handler == "" {
		firstRune, size := utf8.DecodeRuneInString(field.Name)
		upperFirst := unicode.ToUpper(firstRune)

		// encode upperFirst to []byte
		methoName := make([]byte, 4, len(field.Name))
		number := utf8.EncodeRune(methoName, upperFirst)
		methoName = methoName[:number]
		methoName = append(methoName, field.Name[size:]...)

		// assign handler name
		handler = string(methoName)
	}

	// build a TagInfo instance
	info := &tagInfo{
		tagBase: tagBase{Method: method, Host: host, Path: path, Queries: queries},
		handler: handler,
	}

	return info, nil
}
