package convertor

import (
	"regexp"
	"strings"
	"unicode"
)

type TageeTypeKey string

const (
	String  TageeTypeKey = "String"
	Boolean TageeTypeKey = "Boolean"
	boolean TageeTypeKey = "boolean"
	Integer TageeTypeKey = "Integer"
	Long    TageeTypeKey = "Long"
)

var (
	tageeTypeMap = map[TageeTypeKey]string{
		String:  "string",
		Boolean: "bool",
		boolean: "bool",
		Integer: "int32",
		Long:    "int64",
	}
)

type Convertor struct {
	_originName  string
	_originType  string
	_goFieldName string
	_goType      *string
}

func (c *Convertor) GetGoFieldName() string {
	return c._goFieldName
}

func (c *Convertor) GetGoType() string {
	return *c._goType
}

func NewConvertor(tageeName string, tageeType string) *Convertor {
	ist := &Convertor{
		_originName:  tageeName,
		_originType:  tageeType,
		_goFieldName: tageeName,
		_goType:      nil,
	}

	if isPascalCase(tageeName) || isCamelCase(tageeName) {
		ist._goFieldName = capitalizeCamelCase(tageeName)
	} else {
		ist._goFieldName = toUpperCase(tageeName)
	}

	if mapValue, exist := tageeTypeMap[TageeTypeKey(tageeType)]; exist {
		ist._goType = &mapValue
	}

	ist.setIsList(tageeType)

	if ist._goType == nil {
		parts := strings.Split(tageeType, ".")
		str := parts[len(parts)-1]
		ist._goType = &str
	}
	return ist
}

func (c *Convertor) setIsList(tageeType string) {
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
		listStr := "[]"
		mapValue, exist := tageeTypeMap[TageeTypeKey(inner)]
		if exist {
			str := listStr + mapValue
			c._goType = &str
		} else {
			parts := strings.Split(inner, ".")
			str := parts[len(parts)-1]
			if isPascalCase(inner) || isCamelCase(inner) {
				str = listStr + capitalizeCamelCase(inner)
			} else {
				str = listStr + toUpperCase(inner)
			}
			c._goType = &str
		}
	}
}

// utils
func isPascalCase(s string) bool {
	re := regexp.MustCompile(`^[A-Z][a-zA-Z0-9]*$`)
	return re.MatchString(s)
}

func isCamelCase(s string) bool {
	re := regexp.MustCompile(`^[a-z][a-zA-Z0-9]*$`)
	return re.MatchString(s)
}

func capitalizeCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}
