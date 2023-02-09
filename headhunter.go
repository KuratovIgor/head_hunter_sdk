package head_hunter_sdk

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strconv"
)

const (
	authorizeURL = "https://hh.ru/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s"

	authorizeEndpoint = "https://hh.ru/oauth/token"
)

type (
	AuthorizeResponse struct {
		AccessToken  string
		RefreshToken string
	}
)

type Client struct {
	clientID     string
	clientSecret string
	redirectURI  string
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
	}, nil
}

func (c *Client) GetAuthorizationURL(chatID int64) (string, error) {
	fullRedirectURL := c.redirectURI + "/?chat_id=" + strconv.Itoa(int(chatID))

	return fmt.Sprintf(authorizeURL, c.clientID, fullRedirectURL), nil
}

func (c *Client) Authorize(chatID int64, authCode string) (*AuthorizeResponse, error) {
	urlParameters := "grant_type=authorization_code&client_id=" + c.clientID + "&client_secret=" + c.clientSecret + "&redirect_uri=" + c.redirectURI + "/?chat_id=" + strconv.Itoa(int(chatID)) + "&code=" + authCode
	postData := []byte(urlParameters)

	res, reqError := http.Post(authorizeEndpoint, "application/x-www-form-urlencoded", bytes.NewBuffer(postData))

	if reqError != nil {
		return nil, reqError
	}

	defer res.Body.Close()
	data, readError := io.ReadAll(res.Body)

	if readError != nil {
		return nil, readError
	}

	return &AuthorizeResponse{
		AccessToken:  gjson.Get(string(data), "access_token").String(),
		RefreshToken: gjson.Get(string(data), "refreshToken").String(),
	}, nil
}
