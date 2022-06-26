package interfaces

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type checkBalance struct {
	Error   bool
	Message string
	Data    checkBalanceDataBody
}

type checkBalanceDataBody struct {
	Balance  string
	Currency string
}

func CheckBalance() (*checkBalance, int, error) {
	client := NewHttpClient()
	url := "https://bingpay.ng/api/v1/self/balance"
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

	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response checkBalance
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
