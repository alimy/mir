package mir

import (
	"fmt"
	"reflect"
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

// tagMir indicate mir tag common information in struct tag string exclude handler
type tagMir struct {
	Group   string   // Group indicate group information in struct tag string
	Method  string   // Method indicate method information in struct tag string
	Host    string   // Host indicate host information in struct tag string
	Path    string   // Path indicate path information in struct tag string
	Queries []string // Queries indicate path information in struct tag string
}

// tagInfo indicate mir tag information in struct tag string
type tagInfo struct {
	tagMir
	Handler string // Handler indicate handler information in struct tag string
}

// TagField indicate mir tag info used to register route info to engine
type TagField struct {
	tagMir
	Handler interface{} // Handler indicate handler method
}

// TagFields contains *TagField in entry
type TagFields []*TagField

// isGroup return whether mir tag is a group field
func (m *tagInfo) isGroup() bool {
	return m.Group != "" && m.Method == "" && m.Path == ""
}

// TagFieldsFrom build TagFields from entry
func TagFieldsFrom(entry interface{}) (TagFields, error) {
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

	// group information
	var group string

	// get tagFields from entryType and entryPtrType
	tagFields := make(TagFields, 0)
	for i := 0; i < entryType.NumField(); i++ {
		field := entryType.Field(i)
		if tagInfo, err := tagInfoFrom(field); err == nil || err == tagNotExist {
			// group field so just parse group info.group info only have one field
			if tagInfo.isGroup() {
				if group == "" {
					group = tagInfo.Group
					continue
				} else {
					return nil, tagMultGroupInfo
				}
			}
			// method field build tagField first
			if tagField, err := tagFieldFrom(entryValue, entryPtrValue, tagInfo); err == nil {
				tagFields = append(tagFields, tagField)
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// setup group info to tagFields item
	if group != "" {
		for _, tagField := range tagFields {
			tagField.Group = group
		}
	}

	return tagFields, nil
}

// tagFieldFrom build tagField from entry and tagInfo
func tagFieldFrom(entryValue reflect.Value, entryPtrValue reflect.Value, tagInfo *tagInfo) (*TagField, error) {
	var m reflect.Value
	if v := entryValue.MethodByName(tagInfo.Handler); v.IsValid() {
		m = v
	} else if v := entryPtrValue.MethodByName(tagInfo.Handler); v.IsValid() {
		m = v
	} else {
		typeName := entryValue.Type().Name()
		err := fmt.Errorf("not found method %s in struct type %s or *%s", tagInfo.Handler, typeName, typeName)
		return nil, err
	}
	tagField := &TagField{
		tagMir:  tagInfo.tagMir,
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
		return &tagInfo{tagMir: tagMir{Group: tag}}, nil
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
		switch tag[i] {
		case '#':
			i := 1
			for i < len(tag) && tag[i] != '?' {
				i++
			}
			handler = tag[1:i]
			tag = tag[i:]
		case '?':
			i := 1
			j := 1
			queries = make([]string, 0)
			for i < len(tag) && tag[i] != '#' {
				if tag[i] == '&' {
					queries = append(queries, tag[j:i])
					j = i + 1
				}
				i++
			}
			tag = tag[i:]
		}
	}

	// check handler
	if handler == "" {
		firstRune, size := utf8.DecodeRuneInString(field.Name)
		upperFirst := unicode.ToUpper(firstRune)
		methoName := make([]byte, 0, len(field.Name))
		utf8.EncodeRune(methoName, upperFirst)
		methoName = append(methoName, field.Name[size:]...)
		handler = string(methoName)
	}

	return &tagInfo{tagMir: tagMir{Method: method, Host: host, Path: path, Queries: queries}, Handler: handler}, nil
}
