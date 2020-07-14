package besustation

import (
	"os"
	"testing"
)

func TestStationCreationListDeletion(t *testing.T) {
	client := NewClient(os.Getenv("FREIGHTLAYER_API"), os.Getenv("FREIGHTLAYER_API_KEY"))
	station := NewStation("testStation", "test description", "single-org")
	res, err := client.CreateStation(&station)
	if res.StatusCode() != 201 {
		t.Fatalf("Could not create station status code: %d.", res.StatusCode())
	}
	if err != nil {
		t.Fatal(err)
	}

	var station2 Station
	res, err = client.GetStation(station.Id, &station2)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != 200 {
		t.Fatalf("Unable to fetch station %s response was: %d.", station.Id, res.StatusCode())
	}

	if station.Id != station2.Id {
		t.Fatalf("Fetched station id mismatch: expected %s found %s", station.Id, station2.Id)
	}

	var stationgate []Station
	_, err = client.ListStation(&stationgate)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	// @dev TODO: Checks
	countNew := 0
	for _, x := range stationgate {
		t.Logf("\n%v", x)
		if x.Name == "testStation" && (x.State != DELETED && x.State != DELETE_PENDING) {
			res, err = client.DeleteStation(x.Id)
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode() != 202 {
				t.Errorf("Station Deletion Failed Status %d.", res.StatusCode())
			}
			countNew += 1
			t.Logf("\nNew Station: %v", x)
		}
	}
}
