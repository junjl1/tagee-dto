package gen

import (
	"github.com/dave/jennifer/jen"
)

type Field struct {
	structName string
	fieldName  string
	fieldType  string
	comment    string
}

type Generator struct {
	f       *jen.File
	structs map[string][]jen.Code
}

func NewGenerator(pkgName string) *Generator {
	return &Generator{
		f:       jen.NewFile(pkgName),
		structs: make(map[string][]jen.Code),
	}
}

func (g *Generator) SaveFile(fileName string) error {
	return g.f.Save(fileName)
}

func (g *Generator) GenStruct(structName string) {
	if _, exists := g.structs[structName]; !exists {
		return
	}
	if fields, ok := g.structs[structName]; ok {
		g.f.Type().Id(structName).Struct(fields...)
	}
}

func (g *Generator) AppendField(structName string, fieldName string, jsonName string, fieldType string, comment string, required int) {
	var c jen.Code
	if required == 1 {
		c = jen.Id(fieldName).Id(fieldType).Tag(map[string]string{"json": jsonName}).Comment(comment)
	} else {
		c = jen.Id(fieldName).Op("*").Id(fieldType).Tag(map[string]string{"json": jsonName}).Comment(comment)
	}
	if _, exists := g.structs[structName]; !exists {
		g.structs[structName] = []jen.Code{
			c,
		}
		return
	}
	g.structs[structName] = append(g.structs[structName],
		c)
}
