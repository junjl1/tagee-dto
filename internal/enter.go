package internal

import (
	"fmt"
	"github.com/junjl1/tagee-dto/fetcher"
)

func GenTask(tageeCode string) {
	data, err := fetcher.Fetch(tageeCode)
	if err != nil {
		return
	}
	fmt.Println(data)
}
