package airtime

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type buyAirtime struct {
	Phone   string
	Amount  int
	Network int
}

type buyAirtimeRes struct {
	Error   bool
	Message string
	Amount  string
}

func BuyAirtime(phone string, amount, network_id int) (*buyAirtimeRes, int, error) {
	client := bingpay.NewClient()
	url := "https://bingpay.ng/api/v1/buy-airtime"
	method := "POST"
	token := client.Token

	payload := buyAirtime{}
	payload.Phone = phone
	payload.Amount = amount
	payload.Network = network_id
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
	var response buyAirtimeRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
