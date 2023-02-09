package headhunter

type AreaType struct {
	Name string
	Id   string
}

func (c *Client) GetAllAreas() ([]AreaType, error) {
	return c.clientApi.GetAllAreas()
}
