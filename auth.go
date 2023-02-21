package headhunter

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strconv"
)

const (
	authorizeURL      = "https://hh.ru/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s"
	authorizeEndpoint = "https://hh.ru/oauth/token"
	logoutEndpoint    = "https://api.hh.ru/oauth/token"
)

type AuthorizeResponse struct {
	AccessToken  string
	RefreshToken string
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
		RefreshToken: gjson.Get(string(data), "refresh_token").String(),
	}, nil
}

func (c *Client) Logout(token string) error {
	_, reqError := c.sendRequest(methodDELETE, logoutEndpoint, "", token)
	if reqError != nil {
		return reqError
	}

	return nil
}
