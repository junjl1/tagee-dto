package fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/junjl1/tagee-dto/types"
	"io"
	"net/http"
	"strings"
)

func Fetch(tageeCode string) (*types.ResponseData, error) {
	var urlBuilder strings.Builder
	var url string
	if tageeCode != "" {
		// 947b580d9c07c01497040071c7de2572
		urlBuilder.WriteString("https://apidoc.inshopline.com/api/projectInterface/get?code=")
		urlBuilder.WriteString(tageeCode)
		url = urlBuilder.String()
	} else {
		return nil, fmt.Errorf("tageeCode is empty")
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET request failed: %w", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var resp *types.Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}
	return &resp.Data, nil
}
