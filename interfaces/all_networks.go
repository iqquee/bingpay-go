package interfaces

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type allNetworks struct {
	Error   bool
	Message string
	Data    []allNetworksDataBody
}

type allNetworksDataBody struct {
	ID   string
	Name string
	Note string
}

func AllNetworks() (*allNetworks, int, error) {
	client := NewHttpClient()
	url := "https://bingpay.ng/api/v1/all-networks"
	method := "GET"
	token := client.Token
	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		log.Println(reqErr.Error())
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, respErr := client.Http.Do(req)
	if respErr != nil {
		log.Println(respErr.Error())
		return nil, 0, respErr
	}

	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response allNetworks
	if err := json.Unmarshal(resp_body, &response); err != nil {
		log.Println(err)
		return nil, 0, err
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp_body))
	return &response, resp.StatusCode, nil
}
