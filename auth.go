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
	meEndpoint        = "/me"
)

type (
	AuthorizeResponse struct {
		AccessToken  string
		RefreshToken string
	}

	MeType struct {
		Name       string
		LastName   string
		MiddleName string
		Email      string
		Phone      string
		UserID     string
	}
)

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

func (c *Client) GetInfoAboutMe(token string) (*MeType, error) {
	res, err := c.sendRequest(methodGET, baseURL+meEndpoint, "", token)
	if err != nil {
		return nil, err
	}

	var infoAboutMe = &MeType{}

	infoAboutMe.Name = gjson.Get(res, "first_name").String()
	infoAboutMe.LastName = gjson.Get(res, "last_name").String()
	infoAboutMe.MiddleName = gjson.Get(res, "middle_name").String()
	infoAboutMe.Email = gjson.Get(res, "email").String()
	infoAboutMe.Phone = gjson.Get(res, "phone").String()
	infoAboutMe.UserID = gjson.Get(res, "id").String()

	return infoAboutMe, nil
}

func (c *Client) Logout(token string) error {
	_, reqError := c.sendRequest(methodDELETE, logoutEndpoint, "", token)
	if reqError != nil {
		return reqError
	}

	return nil
}
