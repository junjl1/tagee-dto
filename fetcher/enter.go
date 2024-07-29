package fetcher

import (
	"encoding/json"
	"fmt"
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
		urlBuilder.WriteString("https://apidoc.inshopline.com/api/projectInterface/get?code=")
		urlBuilder.WriteString(tageeCode)
		url = urlBuilder.String()
	}
	response, err := http.Get(url)
	//response, err := http.Get(`https://apidoc.inshopline.com/api/projectInterface/get?code=947b580d9c07c01497040071c7de2572`)
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

	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var resp Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	fmt.Println("Response:", resp)
	return
}
