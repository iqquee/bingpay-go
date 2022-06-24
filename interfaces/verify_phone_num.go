package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type VerifyPhoneNum struct {
	Country string
	Phone   string
}

// "country":"NG",
// "phone":"07061785183"
func (vp *VerifyPhoneNum) VerifyPhoneNumber(country, number string) {
	client := NewHttpClient()
	url := ""
	method := "POST"
	token := client.Token

	// payload := VerifyPhoneNum{}
	// payload.Country =
	vp.Country = country //NG for nigeria
	vp.Phone = number
	jsonReq, jsonErr := json.Marshal(&vp)
	if jsonErr != nil {
		log.Println(jsonErr)
		return
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if reqErr != nil {
		log.Println(reqErr)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, respErr := client.Http.Do(req)
	if respErr != nil {
		log.Println(respErr.Error())
		return
	}

	resp_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp_body))
}
