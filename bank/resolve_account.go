package bank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type resloveAccount struct {
	Bank_Code      string
	Account_Number string
}

type resloveAccountRes struct {
	Error bool
	Data  resloveAccountResDataBody
}
type resloveAccountResDataBody struct {
	AccountName   string
	AccountNumber string
	BankCode      string
	ResponseCode  string
	Message       string
}

func ResolveAccount(bank_code, acct_num string) (*resloveAccountRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/resolve-account", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := resloveAccount{}
	payload.Bank_Code = bank_code
	payload.Account_Number = acct_num

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
	var response resloveAccountRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
