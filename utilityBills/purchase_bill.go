package utilitybills

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type purchaseBill struct {
	Service_Id  string
	Customer_Id string
	Variation   string
	Amount      string
}

type purchaseBillRes struct {
	Error          bool
	Message        string
	Data           purchaseBillResDataBody
	Purchased_Code interface{}
}

type purchaseBillResDataBody struct {
	Status          string
	Product_Name    string
	Unique_Element  string
	Unit_Price      int
	Quantity        int
	Channel         string
	Type            string
	Phone           string
	Name            interface{}
	Convinience_fee int
	Amount          int
	Platform        string
}

func PurchaseBill(service_Id, customer_Id, variation, amount string) (*purchaseBillRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/purchase-bill", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := purchaseBill{}
	payload.Service_Id = service_Id
	payload.Customer_Id = customer_Id
	payload.Variation = variation
	payload.Amount = amount

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
	var response purchaseBillRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
