package interfaces

import (
	"fmt"
	"net/http"
)

type HttpClient struct {
	Http  http.Client
	Token string
}

var Token string

func ApiToken(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

func NewHttpClient() *HttpClient {
	var ht HttpClient
	token := ApiToken(Token)
	ht.Http = http.Client{}
	ht.Token = token
	return &HttpClient{
		Http:  ht.Http,
		Token: ht.Token,
	}
}
