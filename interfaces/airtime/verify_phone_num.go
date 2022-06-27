package airtime

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go/interfaces"
)

type verifyPhoneNum struct {
	Country string
	Phone   string
}

type verifyPhoneNumRes struct {
	Error   bool
	Message string
	Data    []verifyPhoneNumResDataBody
}

type verifyPhoneNumResDataBody struct {
	Mobile  string
	Country string
	Name    string
	Status  string
}

func VerifyPhoneNumber(country, number string) (*verifyPhoneNumRes, int, error) {
	client := interfaces.NewHttpClient()
	url := "https://bingpay.ng/api/v1/verify/phone"
	method := "POST"
	token := client.Token

	payload := verifyPhoneNum{}
	payload.Country = country //NG for nigeria
	payload.Phone = number
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
	var response verifyPhoneNumRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}
	return &response, resp.StatusCode, nil
}
