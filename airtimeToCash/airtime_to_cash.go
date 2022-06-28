package airtimetocash

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type airtimeToCash struct {
	Amount  int
	Network int
	Phone   string
}

type airtimeToCashRes struct {
	Error   bool
	Message string
	Data    airtimeToCashResDataBody
}

type airtimeToCashResDataBody struct {
	Amount    string
	Value     int
	Reference int
}

func AirtimeToCash(amount, network_id int, phone string) (*airtimeToCashRes, int, error) {
	client := bingpay.NewClient()
	url := "https://bingpay.ng/api/v1/airtime-cash/process"
	method := "POST"
	token := client.Token

	payload := airtimeToCash{}
	payload.Amount = amount
	payload.Network = network_id
	payload.Phone = phone

	jsonReq, jsonErr := json.Marshal(&payload)
	if jsonErr != nil {
		return nil, 0, jsonErr
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
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
	var response airtimeToCashRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
