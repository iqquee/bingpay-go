package utilitybills

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type allServicesRes struct {
	Error   bool
	Message string
	Data    []allServicesResDataBody
}

type allServicesResDataBody struct {
	Id          string
	Name        string
	Image_Url   string
	Description string
}

func AllServices() (*allServicesRes, int, error) {
	client := bingpay.NewClient()
	url := "https://bingpay.ng/api/v1/all-services"
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
	var response allServicesRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
