package headhunter

import (
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const (
	resumesEndpoint = "/resumes/mine"
)

func (c *Client) GetResumesIds() []string {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", baseURL+resumesEndpoint, nil)
	req.Header.Set("Authorization", "Bearer "+c.token)

	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var resumesIds []string

	value := gjson.Get(string(body), "items")
	for _, item := range value.Array() {
		resumesIds = append(resumesIds, gjson.Get(item.String(), "id").String())
	}

	return resumesIds
}
