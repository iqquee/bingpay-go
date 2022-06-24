package interfaces

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func AllNetworks() {
	client := NewHttpClient()
	url := "https://bingpay.ng/api/v1/all-networks"
	method := "GET"
	token := client.Token
	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		log.Println(reqErr.Error())
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
