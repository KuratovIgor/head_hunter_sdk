package api

import (
	"bytes"
	"github.com/KuratovIgor/head_hunter_sdk"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const (
	authorizeEndpoint = "https://hh.ru/oauth/token"
)

func (api *ClientApi) Authorize(data []byte) (*headhunter.AuthorizeResponse, error) {
	res, reqError := http.Post(authorizeEndpoint, "application/x-www-form-urlencoded", bytes.NewBuffer(data))
	if reqError != nil {
		return nil, reqError
	}

	defer res.Body.Close()

	data, readError := io.ReadAll(res.Body)
	if readError != nil {
		return nil, readError
	}

	return &headhunter.AuthorizeResponse{
		AccessToken:  gjson.Get(string(data), "access_token").String(),
		RefreshToken: gjson.Get(string(data), "refreshToken").String(),
	}, nil
}
