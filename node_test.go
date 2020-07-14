package besustation

import (
	"os"
	"testing"
)

func TestNodeCreation(t *testing.T) {
	client := NewClient(os.Getenv("FREIGHTLAYER_API"), os.Getenv("FREIGHTLAYER_API_KEY"))
	station := NewStation("nodeTestStation", "node creation", "single-org")
	res, err := client.CreateStation(&station)
	t.Logf("%v", station)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Unable to create station status fault: %d.", res.StatusCode())
	}
	defer client.DeleteStation(station.Id)
	env := NewEnvironment("nodeCreate", "just create some nodes", "quorum", "raft")

	res, err = client.CreateEnvironment(station.Id, &env)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Unable to create environment status fault: %d, %s", res.StatusCode(), string(res.Body()))
	}

	var members []Membership
	res, err = client.ListMemberships(station.Id, &members)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Unable to list memberships: %d", res.StatusCode())
	}

	if len(members) != 1 {
		t.Fatalf("Environment unexpected had %d members.", len(members))
	}
	t.Logf("%v", members)

	var nodes []Node
	res, err = client.ListNodes(station.Id, env.Id, &nodes)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode() != 200 {
		t.Fatalf("Unable to list nodes: %d", res.StatusCode())
	}

	t.Logf("Nodes: %v", nodes)

	node := NewNode("testNode", members[0].Id)

	res, err = client.CreateNode(station.Id, env.Id, &node)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 201 {
		t.Fatalf("Creating node failed status fault: %d", res.StatusCode())
	}

	nodes = nil
	res, err = client.ListNodes(station.Id, env.Id, &nodes)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Unable to read nodes status fault: %d", res.StatusCode())
	}

	nodeCount := 2 // @user TODO: we use two nodes so you have a healthcheck 
	if len(nodes) != nodeCount {
		t.Fatalf("Warning, unexpected number of nodes: %d should be %d.", len(nodes), nodeCount)
	}

}
