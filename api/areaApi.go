package api

import (
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const (
	areasEndpoint = "/areas"
)

func (api *ClientApi) GetAllAreas() ([]headhunter.AreaType, error) {
	var allAreas []headhunter.AreaType

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
			var area headhunter.AreaType

			area.Name = gjson.Get(item.String(), "name").String()
			area.Id = gjson.Get(item.String(), "id").String()

			allAreas = append(allAreas, area)
		}
	}

	return allAreas, nil
}
