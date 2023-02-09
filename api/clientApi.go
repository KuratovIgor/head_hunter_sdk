package api

import headhunter "github.com/KuratovIgor/head_hunter_sdk"

const (
	baseURL = "https://api.hh.ru"
)

type ClientApi struct {
	params *headhunter.Params
}

func NewClientApi(params *headhunter.Params) *ClientApi {
	return &ClientApi{params: params}
}
