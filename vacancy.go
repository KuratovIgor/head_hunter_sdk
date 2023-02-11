package headhunter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const (
	vacanciesEndpoint  = "/vacancies"
	applyToJobEndpoint = "/negotiations"

	applyToJobParams = "?vacancy_id=%s&resume_id=%s&message=%s"
)

type (
	Vacancies struct {
		Items []Vacancy
		Pages int
		Page  int
	}

	Vacancy struct {
		Id             string
		Name           string
		Salary         Salary
		Address        Address
		PublishedAt    string
		Employer       string
		Requirement    string
		Responsibility string
		Schedule       string
		AlternateUrl   string
		Area           string
	}

	Address struct {
		City     string
		Street   string
		Building string
	}

	Salary struct {
		From     string
		To       string
		Currency string
	}
)

func (c *Client) GetVacancies() Vacancies {
	res, _ := http.Get(baseURL + vacanciesEndpoint + c.UrlParams.GetQueryString())
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var vacancies Vacancies

	value := gjson.Get(string(body), "items")
	for _, item := range value.Array() {
		var vacancy Vacancy
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

	return vacancies
}

func (c *Client) ApplyToJob(vacancyId string, resumeId string, message string) error {
	params := fmt.Sprintf(applyToJobParams, vacancyId, resumeId, message)

	_, err := c.sendPostRequest(baseURL+applyToJobEndpoint, params)
	if err != nil {
		return err
	}

	return nil
}
