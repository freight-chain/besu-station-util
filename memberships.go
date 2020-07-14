package besustation

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

const (
	memBasePath = "/station/%s/memberships"
)

type Membership struct {
	OrgName string `json:"org_name"`
	Id      string `json:"_id,omitempty"`
}

func NewMembership(orgName string) Membership {
	return Membership{orgName, ""}
}

func (c *KaleidoClient) ListMemberships(stationId string, resultBox *[]Membership) (*resty.Response, error) {
	path := fmt.Sprintf(memBasePath, stationId)
	return c.Client.R().SetResult(resultBox).Get(path)
}
