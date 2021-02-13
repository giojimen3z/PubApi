package config

import (
	"github.com/go-resty/resty/v2"
)

func GetRestyClient() *resty.Client {

	client := resty.New()

	return &resty.Client{
		HostURL:                client.HostURL,
		QueryParam:             client.QueryParam,
		FormData:               client.FormData,
		Header:                 client.Header,
		UserInfo:               client.UserInfo,
		Token:                  client.Token,
		AuthScheme:             client.AuthScheme,
		Cookies:                client.Cookies,
		Error:                  client.Error,
		Debug:                  client.Debug,
		DisableWarn:            client.DisableWarn,
		AllowGetMethodPayload:  client.AllowGetMethodPayload,
		RetryCount:             client.RetryCount,
		RetryWaitTime:          client.RetryWaitTime,
		RetryMaxWaitTime:       client.RetryMaxWaitTime,
		RetryConditions:        client.RetryConditions,
		RetryAfter:             client.RetryAfter,
		JSONMarshal:            client.JSONMarshal,
		JSONUnmarshal:          client.JSONUnmarshal,
		HeaderAuthorizationKey: client.HeaderAuthorizationKey,
	}

}
