package headhunter

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.hh.ru"

	methodGET  = "GET"
	methodPOST = "POST"
)

type Client struct {
	httpClient   *http.Client
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
		httpClient:   &http.Client{},
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		UrlParams:    &Params{},
	}, nil
}

func (c *Client) sendRequest(method string, endpoint string, params string, token string) (string, error) {
	req, err := http.NewRequest(method, endpoint+params, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		err := fmt.Sprintf("API Error! Status code %d", res.StatusCode)
		return "", errors.New(err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
