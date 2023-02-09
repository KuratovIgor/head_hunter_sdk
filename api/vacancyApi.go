package api

import (
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const (
	vacanciesEndpoint = "/vacancies"
)

func (api *ClientApi) GetVacancies(params *headhunter.Params) (*headhunter.Vacancies, error) {
	res, reqError := http.Get(baseURL + vacanciesEndpoint + params.GetQueryString())
	if reqError != nil {
		return nil, reqError
	}

	defer res.Body.Close()

	body, readError := io.ReadAll(res.Body)
	if readError != nil {
		return nil, readError
	}

	var vacancies *headhunter.Vacancies

	value := gjson.Get(string(body), "items")
	for _, item := range value.Array() {
		var vacancy headhunter.Vacancy
		vacancy.Id = gjson.Get(item.String(), "id").String()
		vacancy.Name = gjson.Get(item.String(), "name").String()
		vacancy.Salary.From = gjson.Get(item.String(), "salary.from").String()
		vacancy.Salary.To = gjson.Get(item.String(), "salary.to").String()
		vacancy.Salary.Currency = gjson.Get(item.String(), "salary.currency").String()
		vacancy.Address.City = gjson.Get(item.String(), "address.city").String()
		vacancy.Address.Street = gjson.Get(item.String(), "address.street").String()
		vacancy.Address.Building = gjson.Get(item.String(), "address.building").String()
		vacancy.PublishedAt = gjson.Get(item.String(), "published_at").String()
		vacancy.Employer = gjson.Get(item.String(), "employer.name").String()
		vacancy.Requirement = gjson.Get(item.String(), "snippet.requirement").String()
		vacancy.Responsibility = gjson.Get(item.String(), "snippet.responsibility").String()
		vacancy.Schedule = gjson.Get(item.String(), "schedule.name").String()
		vacancy.AlternateUrl = gjson.Get(item.String(), "alternate_url").String()
		vacancy.Area = gjson.Get(item.String(), "area.name").String()

		vacancies.Items = append(vacancies.Items, vacancy)
	}

	return vacancies, nil
}
