package interfaces

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CheckBalance(bearerToken string) {
	client := http.Client{}
	url := "https://bingpay.ng/api/v1/self/balance"
	method := "GET"
	token := fmt.Sprintf("Bearer %s", bearerToken)
	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		log.Println(reqErr.Error())
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Println(respErr.Error())
		return
	}

	resp_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp_body))
}
