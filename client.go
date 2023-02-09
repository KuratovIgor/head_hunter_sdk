package headhunter

import (
	"errors"
)

const (
	baseURL = "https://api.hh.ru"
)

type Client struct {
	clientID     string
	clientSecret string
	redirectURI  string
	UrlParams    *Params
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

	return &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		UrlParams:    NewParams(),
	}, nil
}
