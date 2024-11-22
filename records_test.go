package netacl

import (
	"log"
	"testing"
)

func TestCreateNewSRVRecord(t *testing.T) {
	setTestKey(t)
	if err := c.NewRecords("mcsh.io", &SRVRecords{
		{
			Name:     "_minecraft._tcp.extra.mcsh.io",
			Priority: 0,
			Weight:   5,
			Port:     25565,
			Target:   "bth01-fra.iona.sh",
		},
	}); err != nil {
		t.Error(err)
	}
}

func TestCreateDuplicateSRVRecord(t *testing.T) {
	setTestKey(t)
	if err := c.NewRecords("mcsh.io", &SRVRecords{
		{
			Name:     "_minecraft._tcp.roolps.mcsh.io",
			Priority: 0,
			Weight:   5,
			Port:     25565,
			Target:   "bth01-fra.iona.sh",
		},
	}); err != nil {
		if err != ErrRecordExists {
			t.Error(err)
		}
	} else {
		t.Error("created duplicate record, expected error")
	}
}

func TestGetBodyFromCreateSRVRecordRequest(t *testing.T) {
	setTestKey(t)
	recs := &SRVRecords{
		{
			Name:     "_minecraft._tcp.roolps-test-3.mcsh.io",
			Priority: 0,
			Weight:   5,
			Port:     25566,
			Target:   "bth01-fra.iona.sh",
		},
	}

	if err := c.NewRecords("mcsh.io", recs); err != nil {
		t.Error(err)
	}
	if recs == nil {
		t.Error("no records returned, expected 1")
	}
	r := *recs
	if len(r) == 0 {
		t.Error("no records returned, expected 1")
	}
	log.Println(r[0].ID)
}
