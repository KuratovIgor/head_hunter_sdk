package headhunter

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const (
	resumesEndpoint = "/resumes/mine"

	nameString   = "%s %s %s"
	salaryString = "%s %s"
)

type (
	Resume struct {
		Id        string
		Name      string
		Title     string
		Area      string
		Age       string
		Salary    string
		Download  string
		URL       string
		Education string
	}
)

func (c *Client) GetResumesIds() ([]string, error) {
	var resumesIds []string

	res, err := c.sendRequest(methodGET, baseURL+resumesEndpoint, "")
	if err != nil {
		return nil, err
	}

	items := gjson.Get(res, "items")
	for _, item := range items.Array() {
		resumesId := gjson.Get(item.String(), "id").String()
		resumesIds = append(resumesIds, resumesId)
	}

	return resumesIds, nil
}

func (c *Client) GetResumes() ([]Resume, error) {
	var resumes []Resume

	res, err := c.sendRequest(methodGET, baseURL+resumesEndpoint, "")
	if err != nil {
		return nil, err
	}

	items := gjson.Get(res, "items")
	for _, item := range items.Array() {
		var resume Resume

		resume.Id = gjson.Get(item.String(), "id").String()
		resume.Name = fmt.Sprintf(nameString, gjson.Get(item.String(), "last_name").String(), gjson.Get(item.String(), "first_name").String(), gjson.Get(item.String(), "middle_name").String())
		resume.Title = gjson.Get(item.String(), "title").String()
		resume.Area = gjson.Get(item.String(), "area.name").String()
		resume.Age = gjson.Get(item.String(), "age").String()
		resume.Salary = fmt.Sprintf(salaryString, gjson.Get(item.String(), "salary.amount").String(), gjson.Get(item.String(), "salary.currency").String())
		resume.Download = gjson.Get(item.String(), "actions.download.pdf.url").String()
		resume.URL = gjson.Get(item.String(), "alternate_url").String()
		resume.Education = gjson.Get(item.String(), "education.level.name").String()

		resumes = append(resumes, resume)
	}

	return resumes, nil
}
