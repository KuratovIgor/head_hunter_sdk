package headhunter

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
