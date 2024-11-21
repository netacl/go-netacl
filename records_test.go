package netacl

import "testing"

func TestCreateNewSRVRecord(t *testing.T) {
	setTestKey(t)
	if err := c.NewRecords("", &SRVRecords{
		{
			Name:     "",
			Priority: 0,
			Weight:   5,
			Port:     25565,
			Target:   "",
		},
	}); err != nil {
		t.Error(err)
	}
}
