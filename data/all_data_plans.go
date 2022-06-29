package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type allDataPlans struct {
	Error   bool
	Message string
	Data    []allDataPlansDataBody
}

type allDataPlansDataBody struct {
	Id         string
	Network_Id string
	Name       string
	Price      string
	Uniq_Id    string
}

func AllDataPlans() (*allDataPlans, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/all-data-plans", client.BaseUrl)
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
	var response allDataPlans
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
