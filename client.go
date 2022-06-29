package bingpay

import (
	"fmt"
	"net/http"
)

type Client struct {
	Http    http.Client
	Token   string
	BaseUrl string
}

var Token string

func ApiToken(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

func NewClient() *Client {
	var cl Client
	token := ApiToken(Token)
	cl.Http = http.Client{}
	cl.Token = token
	cl.BaseUrl = "https://bingpay.ng/api/v1"
	return &Client{
		Http:    cl.Http,
		Token:   cl.Token,
		BaseUrl: cl.BaseUrl,
	}
}
