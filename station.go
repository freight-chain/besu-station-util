package besustation

import (
	"fmt"

	"gopkg.in/resty.v1"
)

const (
	SingleOrg      = "single-org"
	MultiOrg       = "multi-org"
	DELETE_PENDING = "delete_pending"
	DELETED        = "deleted"
)

type BesuClient struct {
	Client *resty.Client
}

type Station struct {
	Id          string `json:"_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Mode        string `json:"mode"`
	DeletedAt   string `json:"deleted_at,omitempty"`
	State       string `json:"state,omitempty"`
}

func NewStation(name, description, mode string) Station {
	return Station{
		Id:          "",
		Name:        name,
		Description: description,
		Mode:        mode,
		DeletedAt:   "",
		State:       "",
	}
}

func NewClient(api string, apiKey string) BesuClient {
	r := resty.New().SetHostURL(api).SetAuthToken(apiKey)
	return BesuClient{r}
}

func (c *BesuClient) CreateStation(station *Station) (*resty.Response, error) {
	return c.Client.R().SetBody(station).SetResult(station).Post("/station")
}

func (c *BesuClient) ListStation(resultBox *[]Station) (*resty.Response, error) {
	return c.Client.R().SetResult(resultBox).Get("/station")
}

func (c *BesuClient) GetStation(id string, resultBox *Station) (*resty.Response, error) {
	path := fmt.Sprintf("/station/%s", id)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *BesuClient) DeleteStation(stationId string) (*resty.Response, error) {
	return c.Client.R().Delete(fmt.Sprintf("/station/%s", stationId))
}
