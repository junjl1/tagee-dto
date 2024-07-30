package enum

import (
	"regexp"
	"strings"
)

type TageeTypeKey string

const (
	String  TageeTypeKey = "String"
	Boolean TageeTypeKey = "Boolean"
	Integer TageeTypeKey = "Integer"
	Long    TageeTypeKey = "Long"
)

var (
	tageeTypeMap = map[TageeTypeKey]string{
		String:  "string",
		Boolean: "bool",
		Integer: "int32",
		Long:    "int64",
	}
)

type TageeType struct {
	_origin string
	_goType *string
}

func (t *TageeType) GetGoType() string {
	return *t._goType
}

func NewTageeType(tageeType string) *TageeType {
	ist := &TageeType{
		_origin: tageeType,
	}
	if mapValue, exist := tageeTypeMap[TageeTypeKey(tageeType)]; exist {
		ist._goType = &mapValue
	}

	ist.setIsList(tageeType)

	if ist._goType == nil {
		str := strings.ToUpper(tageeType)
		ist._goType = &str
	}
	return ist
}

func (t *TageeType) setIsList(tageeType string) {
	pattern := `List<(.*?)>`
	// 编译正则表达式
	res, err := regexp.Compile(pattern)
	if err != nil {
		// 正则表达式编译错误
		return
	}
	match := res.FindStringSubmatch(tageeType)
	if len(match) == 0 {
		return
	}
	wrapper := match[0]
	inner := match[1]
	if wrapper != "" && inner != "" {
		str := "[]"
		mapValue, exist := tageeTypeMap[TageeTypeKey(inner)]
		if exist {
			str := str + mapValue
			t._goType = &str
		} else {
			str := str + strings.ToUpper(inner)
			t._goType = &str
		}
	}
}
