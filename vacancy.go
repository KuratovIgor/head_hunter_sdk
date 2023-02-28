package headhunter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"net/url"
)

const (
	vacancyEndpoint    = "/vacancies/%s"
	vacanciesEndpoint  = "/vacancies"
	applyToJobEndpoint = "/negotiations"

	applyToJobParams = "?vacancy_id=%s&resume_id=%s&message=%s"
)

type (
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

	Response struct {
		State   string
		Vacancy ShortVacancy
	}

	ShortVacancy struct {
		Name         string
		Salary       Salary
		Employer     string
		AlternateUrl string
		Area         string
	}
)

func (c *Client) GetVacancies(token string) ([]Vacancy, error) {
	res, err := c.sendRequest(methodGET, baseURL+vacanciesEndpoint, c.UrlParams.GetQueryString(), token)
	if err != nil {
		return nil, err
	}
	var vacancies []Vacancy

	items := gjson.Get(res, "items")
	for _, item := range items.Array() {
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

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func (c *Client) GetVacancy(vacancyId string, token string) (*Vacancy, error) {
	res, err := c.sendRequest(methodGET, baseURL+fmt.Sprintf(vacancyEndpoint, vacancyId), c.UrlParams.GetQueryString(), token)
	if err != nil {
		return nil, err
	}
	var vacancy = &Vacancy{}

	vacancy.Id = gjson.Get(res, "id").String()
	vacancy.Name = gjson.Get(res, "name").String()
	vacancy.Salary.From = gjson.Get(res, "salary.from").String()
	vacancy.Salary.To = gjson.Get(res, "salary.to").String()
	vacancy.Salary.Currency = gjson.Get(res, "salary.currency").String()
	vacancy.Address.City = gjson.Get(res, "address.city").String()
	vacancy.Address.Street = gjson.Get(res, "address.street").String()
	vacancy.Address.Building = gjson.Get(res, "address.building").String()
	vacancy.PublishedAt = gjson.Get(res, "published_at").String()
	vacancy.Employer = gjson.Get(res, "employer.name").String()
	vacancy.Requirement = gjson.Get(res, "snippet.requirement").String()
	vacancy.Responsibility = gjson.Get(res, "snippet.responsibility").String()
	vacancy.Schedule = gjson.Get(res, "schedule.name").String()
	vacancy.AlternateUrl = gjson.Get(res, "alternate_url").String()
	vacancy.Area = gjson.Get(res, "area.name").String()

	return vacancy, nil
}

func (c *Client) ApplyToJob(vacancyId string, resumeId string, message string, token string) error {
	params := fmt.Sprintf(applyToJobParams, vacancyId, resumeId, url.QueryEscape(message))

	_, err := c.sendRequest(methodPOST, baseURL+applyToJobEndpoint, params, token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetResponseList(token string) ([]Response, error) {
	res, err := c.sendRequest(methodGET, baseURL+applyToJobEndpoint, "", token)
	if err != nil {
		return nil, err
	}

	var responses []Response

	items := gjson.Get(res, "items")
	for _, item := range items.Array() {
		if len(responses) == 5 {
			break
		}

		var response Response
		var vacancy ShortVacancy

		response.State = gjson.Get(item.String(), "state.name").String()

		vacancy.Name = gjson.Get(item.String(), "vacancy.name").String()
		vacancy.Salary.From = gjson.Get(item.String(), "vacancy.salary.from").String()
		vacancy.Salary.To = gjson.Get(item.String(), "vacancy.salary.to").String()
		vacancy.Salary.Currency = gjson.Get(item.String(), "vacancy.salary.currency").String()
		vacancy.Employer = gjson.Get(item.String(), "vacancy.employer.name").String()
		vacancy.AlternateUrl = gjson.Get(item.String(), "vacancy.alternate_url").String()
		vacancy.Area = gjson.Get(item.String(), "vacancy.area.name").String()

		response.Vacancy = vacancy
		responses = append(responses, response)
	}

	return responses, nil
}
