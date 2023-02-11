package headhunter

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.hh.ru"
)

type Client struct {
	httpClient   *http.Client
	token        string
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
		UrlParams:    NewParams(),
	}, nil
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) IsTokenExists() bool {
	if c.token == "" {
		return false
	}

	return true
}

func (c *Client) sendPostRequest(endpoint string, params string) (url.Values, error) {
	req, err := http.NewRequest("POST", endpoint+params, nil)
	if err != nil {
		return url.Values{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return url.Values{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		err := fmt.Sprintf("API Error! Status code %d", res.StatusCode)
		return url.Values{}, errors.New(err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return url.Values{}, err
	}

	values, err := url.ParseQuery(string(resBody))
	if err != nil {
		return url.Values{}, err
	}

	return values, nil
}
