package bank

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type allBanksRes struct {
	Error   bool
	Message string
	Data    []allBanksResDataBody
}

type allBanksResDataBody struct {
	Name        string
	Slug        string
	Code        string
	LongCode    string
	Gateway     interface{}
	PayWithBank bool
	Active      bool
	IsDeleted   bool
	Country     string
	Currency    string
	Type        string
	Id          int
	CreatedAt   string
	UpdatedAt   string
}

func AllBanks() (*allBanksRes, int, error) {
	client := bingpay.NewClient()
	url := "https://bingpay.ng/api/v1/all-banks"
	method := "POST"
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
	var response allBanksRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
