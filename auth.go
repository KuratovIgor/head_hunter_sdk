package headhunter

import (
	"fmt"
	"strconv"
)

const authorizeURL = "https://hh.ru/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s"

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

	return c.clientApi.Authorize(postData)
}
