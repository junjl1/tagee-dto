package gen

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

func GenDto() {
	f := jen.NewFile("resDto")

	//
	f.Type().Id("Person").Struct(
		jen.Id("Name").String().Tag(map[string]string{"json": "name"}).Comment("aaaa"),
		jen.Id("Age").Int().Tag(map[string]string{"json": "age"}).Comment("bbbb"),
	)

	err := f.Save("result.go")
	if err != nil {
		fmt.Errorf("failed to save file: %w", err)
	}
}
