package internal

import (
	"fmt"
	"github.com/junjl1/tagee-dto/enum"
	"github.com/junjl1/tagee-dto/fetcher"
	"github.com/junjl1/tagee-dto/gen"
	"strings"
)

func GenTask(tageeCode string) {
	data, err := fetcher.Fetch(tageeCode)
	if err != nil {
		return
	}
	g := gen.NewGenerator("myPkgName")
	for idx, value := range data.InputParam {
		if idx == 0 && value.Key == "root" {
			continue
		}
		parts := strings.Split(value.Key, ".")
		structName := parts[len(parts)-1]
		for _, dto := range value.DetailList {
			t := enum.NewTageeType(dto.Type)
			g.AppendField(structName, dto.Name, t.GetGoType(), dto.Comment)
		}
		g.GenStruct(structName)
	}
	err = g.SaveFile("res.go")
	if err != nil {
		fmt.Println("opps", err)
		return
	}
	//typeIst := enum.NewTageeType("Integer")
	//fmt.Println(typeIst.GetGoType())

	//g := gen.NewGenerator("myPkgName")
	//g.AppendField("Person", "name", "string", "person name")
	//g.AppendField("Person", "age", "int", "person age")
	//g.GenStruct("Person")
	//g.AppendField("person", "name", "string", "person name")
	//g.AppendField("person", "age", "Person", "person age")
	//g.GenStruct("person")
	//fmt.Printf("%#v", g.GetF())
}
