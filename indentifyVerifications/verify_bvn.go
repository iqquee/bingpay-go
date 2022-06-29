package indentifyverifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go"
)

type verifyBvn struct {
	FirstName string
	LastName  string
	Phone     string
	Bvn       string
}

type verifyBvnRes struct {
	Error   bool
	Message string
	Data    verifyBvnResDataBody
}

type verifyBvnResDataBody struct {
	Bvn          string
	FirstName    string
	LastName     string
	MiddleName   string
	Phone        string
	BirthDate    string
	Gender       string
	Nationality  string
	Photo        string
	FieldMatches verifyBvnResDataBodyFieldMatches
}

type verifyBvnResDataBodyFieldMatches struct {
	LastName  bool
	FirstName bool
}

func VerifyBvn(firstName, lastName, phone, bvn string) (*verifyBvnRes, int, error) {
	client := bingpay.NewClient()
	url := fmt.Sprintf("%s/verify/bvn", client.BaseUrl)
	method := "POST"
	token := client.Token

	payload := verifyBvn{}
	payload.FirstName = firstName
	payload.LastName = lastName
	payload.Phone = phone
	payload.Bvn = bvn

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
	var response verifyBvnRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
