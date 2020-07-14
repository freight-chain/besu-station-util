package besustation

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

type Environment struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Provider      string `json:"provider"`
	ConsensusType string `json:"consensus_type"`
	Id            string `json:"_id,omitempty"`
}

const (
	envBasePath = "/station/%s/environments"
)

func NewEnvironment(name, description, provider, consensus string) Environment {
	return Environment{
		Name:          name,
		Description:   description,
		Provider:      provider,
		ConsensusType: consensus,
		Id:            "",
	}
}

func (c *BesuClient) ListEnvironments(stationId string, resultBox *[]Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath, stationId)
	return c.Client.R().SetResult(resultBox).Get(path)
}

func (c *BesuClient) CreateEnvironment(stationId string, environment *Environment) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath, stationId)
	return c.Client.R().SetResult(environment).SetBody(environment).Post(path)
}

func (c *BesuClient) DeleteEnvironment(stationId, environmentId string) (*resty.Response, error) {
	path := fmt.Sprintf(envBasePath+"/%s", stationId, environmentId)
	return c.Client.R().Delete(path)
}
