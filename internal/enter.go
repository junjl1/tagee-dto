package internal

import (
	"fmt"
	"github.com/junjl1/tagee-dto/convertor"
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
	input := data.InputParam
	for idx := len(input) - 1; idx >= 0; idx-- {
		if idx == 0 && input[idx].Key == "root" {
			continue
		}
		parts := strings.Split(input[idx].Key, ".")
		structName := parts[len(parts)-1]
		for _, dto := range input[idx].DetailList {
			t := convertor.NewConvertor(dto.Name, dto.Type)
			g.AppendField(structName, t.GetGoFieldName(), dto.Name, t.GetGoType(), dto.Comment)
		}
		g.GenStruct(structName)
	}
	err = g.SaveFile("res.go")
	if err != nil {
		fmt.Println("opps", err)
		return
	}
}
