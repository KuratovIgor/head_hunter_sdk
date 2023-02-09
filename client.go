package headhunter

import (
	"errors"
	"github.com/KuratovIgor/head_hunter_sdk/api"
)

type Client struct {
	clientID     string
	clientSecret string
	redirectURI  string
	clientApi    *api.ClientApi
}

func NewClient(clientID string, clientSecret string, redirectURI string) (*Client, error) {
	if clientID == "" {
		return nil, errors.New("client id is empty")
	}

	if clientSecret == "" {
		return nil, errors.New("client secret is empty")
	}

	if redirectURI == "" {
		return nil, errors.New("redirect uri is empty")
	}

	clientApi := api.NewClientApi()

	return &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		clientApi:    clientApi,
	}, nil
}
