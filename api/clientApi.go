package api

const (
	baseURL = "https://api.hh.ru"
)

type ClientApi struct {
}

func NewClientApi() *ClientApi {
	return &ClientApi{}
}
