package bingpay

import (
	"fmt"
	"net/http"
)

type Client struct {
	Http  http.Client
	Token string
}

var Token string

func ApiToken(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

func NewClient() *Client {
	var ht Client
	token := ApiToken(Token)
	ht.Http = http.Client{}
	ht.Token = token
	return &Client{
		Http:  ht.Http,
		Token: ht.Token,
	}
}
