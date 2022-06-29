package utilitybills

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type serviceVariationRes struct {
	Error   bool
	Message string
	Data    serviceVariationDataBody
}

type serviceVariationDataBody struct {
	ServiceName     string
	ServiceId       string
	Convinience_Fee string
	Variations      []variationsDataBody
}

type variationsDataBody struct {
	Variation_Code   string
	Name             string
	Variation_Amount string
	FixedPrice       string
}

func ServiceVariation(service_id int) (*serviceVariationRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/service/%d", client.BaseUrl, service_id)
	method := "GET"
	token := client.Token

	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, respErr := client.Http.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response serviceVariationRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
