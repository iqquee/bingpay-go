package indentifyverifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/iqquee/bingpay-go"
)

type verifyBvn struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Bvn       string `json:"bvn"`
}

type verifyBvnRes struct {
	Error   bool                 `json:"error"`
	Message string               `json:"message"`
	Data    verifyBvnResDataBody `json:"data"`
}

type verifyBvnResDataBody struct {
	Bvn          string                           `json:"bvn"`
	FirstName    string                           `json:"first_name"`
	LastName     string                           `json:"last_name"`
	MiddleName   string                           `json:"middle_name"`
	Phone        string                           `json:"phone"`
	BirthDate    string                           `json:"birth_date"`
	Gender       string                           `json:"gender"`
	Nationality  string                           `json:"nationaliy"`
	Photo        string                           `json:"photo"`
	FieldMatches verifyBvnResDataBodyFieldMatches `json:"field_matches"`
}

type verifyBvnResDataBodyFieldMatches struct {
	LastName  bool `json:"last_name"`
	FirstName bool `json:"first_name"`
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
