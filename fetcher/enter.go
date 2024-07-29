package fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/junjl1/tagee-dto/types"
	"io/ioutil"
	"net/http"
	"strings"
)

func Fetch(tageeCode string) {
	var urlBuilder strings.Builder
	var url string
	if tageeCode == "" {
		url = ""
	} else {
		// 947b580d9c07c01497040071c7de2572
		urlBuilder.WriteString("https://apidoc.inshopline.com/api/projectInterface/get?code=")
		urlBuilder.WriteString(tageeCode)
		url = urlBuilder.String()
	}
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	var resp types.Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	fmt.Println("Response:", resp.Data)
	return
}
