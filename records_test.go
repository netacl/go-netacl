package netacl

import (
	"testing"

	"github.com/roolps/logging"
)

func TestGetRecords(t *testing.T) {
	setTestKey(t)
	r := &SRVRecords{}
	if err := c.GetRecords("mcsh.io", r); err != nil {
		t.Error(err)
	}
}

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
	logging.Debug(r[0].ID)
}

func TestDeleteRecordsWithInvalidID(t *testing.T) {
	setTestKey(t)
	if err := c.DeleteRecords("mcsh.io", &SRVRecords{{ID: "testid"}}); err == nil {
		t.Error("expected error, received none")
	} else {
		if err != ErrRecordDoesntExistInZone {
			t.Errorf("unexpected error returned: %v", err)
		}
	}
}

func TestDeleteRecordsWithValidID(t *testing.T) {
	setTestKey(t)
	r := &SRVRecords{}
	if err := c.GetRecords("mcsh.io", r); err != nil {
		t.Errorf("unexpected error returned: %v", err)
	}
	for _, rec := range *r {
		logging.Debug(rec.ID)
	}

	if err := c.DeleteRecords("mcsh.io", &SRVRecords{{ID: "X21pbmVjcmFmdC5fdGNwLnJvb2xwcy10ZXN0LTMubWNzaC5pby9TUlZ8MzEy"}}); err != nil {
		t.Errorf("unexpected error returned: %v", err)
	}
}
