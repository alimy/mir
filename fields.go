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

// tagInfo indicate mir tag information in struct tag string
type tagInfo struct {
	Group   string   // Group indicate group information in struct tag string
	Metod   string   // Method indicate method information in struct tag string
	Host    string   // Host indicate host information in struct tag string
	Path    string   // Path indicate path information in struct tag string
	Queries []string // Queries indicate path information in struct tag string
	Handler string   // Handler indicate handler information in struct tag string
}

// TagField indicate mir tag info used to register route info to engine
type TagField struct {
	Group   string      // Group indicate group information
	Metod   string      // Method indicate method information
	Host    string      // Host indicate host information
	Path    string      // Path indicate path information
	Queries []string    // Queries indicate path information
	Handler interface{} // Handler indicate handler method
}

// TagFields contains *TagField in entry
type TagFields []*TagField

// TagFieldsFrom build TagFields from entry
func TagFieldsFrom(entry interface{}) (TagFields, error) {
	// used to find method from T and lookup struct tag string of mir tag information
	entryType := reflect.TypeOf(entry)

	if entryType.Kind() == reflect.Ptr {
		entryType = entryType.Elem()
	}
	if entryType.Kind() != reflect.Struct {
		return nil, tagNotValideType
	}

	// used to find method from *T
	entryPtrType := reflect.PtrTo(entryType)

	// get tagFields from entryType and entryPtrType
	tagFields := make(TagFields, 0)
	for i := 0; i < entryType.NumField(); i++ {
		field := entryType.Field(i)
		if tagInfo, err := tagInfoFrom(field); err == nil || err == tagNotExist {
			if tagField, err := tagFieldFrom(entryType, entryPtrType, tagInfo); err == nil {
				tagFields = append(tagFields, tagField)
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// check tagFields and set group info if need
	var group string
	for _, tagField := range tagFields {
		if tagField.Group != "" {
			if group == "" {
				group = tagField.Group
			} else {
				return nil, tagMultGroupInfo
			}
		}
	}
	if group != "" {
		for _, tagField := range tagFields {
			tagField.Group = group
		}
	}

	return tagFields, nil
}

// tagFieldFrom build tagField from entry and tagInfo
func tagFieldFrom(entryType reflect.Type, entryPtrType reflect.Type, tagInfo *tagInfo) (tagField *TagField, err error) {
	var m reflect.Method
	if method, ok := entryType.MethodByName(tagInfo.Handler); ok {
		m = method
	} else if method, ok = entryPtrType.MethodByName(tagInfo.Handler); ok {
		m = method
	} else {
		err = fmt.Errorf("not found methd %s in struct %s or *%s", tagInfo.Handler, entryType.Name(), entryType.Name())
	}
	tagField = &TagField{
		Host:    tagInfo.Host,
		Path:    tagInfo.Path,
		Queries: tagInfo.Queries,
		Handler: m.Func.Interface()}
	return
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
		return &tagInfo{Group: tag}, nil
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

	return &tagInfo{Metod: method, Host: host, Path: path, Queries: queries, Handler: handler}, nil
}
