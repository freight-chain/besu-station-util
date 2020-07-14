package besustation

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

type Node struct {
	Name         string `json:"name"`
	MembershipId string `json:"membership_id"`
	Id           string `json:"_id,omitempty"`
}

func NewNode(name, membershipId string) Node {
	return Node{
		Name:         name,
		MembershipId: membershipId,
		Id:           "",
	}
}

func (c *BesuClient) CreateNode(station, envId string, node *Node) (*resty.Response, error) {
	path := fmt.Sprintf("/station/%s/environments/%s/nodes", station, envId)
	return c.Client.R().SetResult(node).SetBody(node).Post(path)
}

func (c *BesuClient) DeleteNode(station, envId, nodeId string) (*resty.Response, error) {
	path := fmt.Sprintf("/station/%s/environments/%s/nodes/%s", station, envId, nodeId)
	return c.Client.R().Delete(path)
}

func (c *BesuClient) ListNodes(station, envId string, resultBox *[]Node) (*resty.Response, error) {
	path := fmt.Sprintf("/station/%s/environments/%s/nodes", station, envId)
	return c.Client.R().SetResult(resultBox).Get(path)
}
