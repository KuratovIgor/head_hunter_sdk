package headhunter

import (
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const (
	areasEndpoint = "/areas"
)

type AreaType struct {
	Name string
	Id   string
}

func (c *Client) GetAllAreas() ([]AreaType, error) {
	var allAreas []AreaType

	res, reqError := http.Get(baseURL + areasEndpoint)
	if reqError != nil {
		return nil, reqError
	}

	defer res.Body.Close()

	body, readError := io.ReadAll(res.Body)
	if readError != nil {
		return nil, readError
	}

	largeAreas := gjson.Parse(string(body))

	for _, largeArea := range largeAreas.Array() {
		areas := gjson.Get(largeArea.String(), "areas")

		for _, item := range areas.Array() {
			var area AreaType

			area.Name = gjson.Get(item.String(), "name").String()
			area.Id = gjson.Get(item.String(), "id").String()

			allAreas = append(allAreas, area)
		}
	}

	return allAreas, nil
}
