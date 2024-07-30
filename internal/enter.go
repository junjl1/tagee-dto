package internal

import (
	"fmt"
	"github.com/junjl1/tagee-dto/convertor"
	"github.com/junjl1/tagee-dto/fetcher"
	"github.com/junjl1/tagee-dto/gen"
	"github.com/junjl1/tagee-dto/types"
	"strings"
)

func GenTask(tageeCode string) {
	data, err := fetcher.Fetch(tageeCode)
	if err != nil {
		return
	}
	genSubTask("dto", "req.go", data.InputParam)
	genSubTask("dto", "res.go", data.OutputParam)
}

func genSubTask(pkgName string, fileName string, data []types.ParamDTO) {
	g := gen.NewGenerator(pkgName)
	for idx := len(data) - 1; idx >= 0; idx-- {
		if idx == 0 && data[idx].Key == "root" {
			continue
		}
		parts := strings.Split(data[idx].Key, ".")
		structName := parts[len(parts)-1]
		for _, dto := range data[idx].DetailList {
			t := convertor.NewConvertor(dto.Name, dto.Type)
			g.AppendField(structName, t.GetGoFieldName(), dto.Name, t.GetGoType(), dto.Comment)
		}
		g.GenStruct(structName)
	}
	err := g.SaveFile(fileName)
	if err != nil {
		fmt.Println("opps", err)
		return
	}
}
