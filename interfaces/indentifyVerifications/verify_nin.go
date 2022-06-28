package indentifyverifications

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hisyntax/bingpay-go/interfaces"
)

type verifyNin struct {
	FirstName string
	LastName  string
	Phone     string
	Nin       string
}

type verifyNinRes struct {
	Error   bool
	Message string
	Data    verifyNinResDataBody
}

type verifyNinResDataBody struct {
	Nin              string
	Title            string
	FirstName        string
	LastName         string
	MiddleName       string
	Phone            string
	BirthDate        string
	Nationality      string
	Gender           string
	Profession       string
	StateOfOrigin    string
	LgaOfOrigin      string
	PlaceOfOrigin    string
	Photo            string
	MaritalStatus    string
	Height           string
	Email            string
	EmploymentStatus string
	BirthState       string
	BirthCountry     string
	NextOfKin        verifyNinResDataBodyNextOfKin
	NspokenLang      string
	OspokenLang      string
	ParentLastName   string
	Religion         string
	Residence        verifyNinResDataBodyResidence
	Signature        string
	FieldMatches     verifyNinResDataBodyFieldMatches
}
type verifyNinResDataBodyNextOfKin struct {
	FirstName  string
	LastName   string
	MiddleName string
	Address1   string
	Address2   string
	Lga        string
	State      string
	Town       string
}

type verifyNinResDataBodyResidence struct {
	Address1 string
	Address2 string
	Town     string
	Lga      string
	State    string
	Status   string
}

type verifyNinResDataBodyFieldMatches struct {
	LastName  bool
	FirstName bool
}

func VerifyNin(firstName, lastName, phone, nin string) (*verifyNinRes, int, error) {
	client := interfaces.NewHttpClient()
	url := "https://bingpay.ng/api/v1/verify/nin"
	method := "POST"
	token := client.Token

	payload := verifyNin{}
	payload.FirstName = firstName
	payload.LastName = lastName
	payload.Phone = phone
	payload.Nin = nin

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
	var response verifyNinRes
	if err := json.Unmarshal(resp_body, &response); err != nil {
		return nil, 0, err
	}

	return &response, resp.StatusCode, nil
}
