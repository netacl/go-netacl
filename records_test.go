package netacl

import (
	"log"
	"testing"
)

func TestGetRecords(t *testing.T) {
	setTestKey(t)
	r := &SRVRecords{}
	if err := c.GetRecords("peanuts.org", r); err != nil {
		t.Error(err)
	}
	log.Println(r)
}

func TestCreateNewSRVRecord(t *testing.T) {
	setTestKey(t)
	if err := c.NewRecords("peanuts.org", &SRVRecords{
		{
			Name:     "_minecraft._tcp.extra.peanuts.org",
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
	if err := c.NewRecords("peanuts.org", &SRVRecords{
		{
			Name:     "_minecraft._tcp.roolps.peanuts.org",
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
			Name:     "_minecraft._tcp.roolps-test-0.peanuts.org",
			Priority: 0,
			Weight:   5,
			Port:     25566,
			Target:   "bth01-fra.iona.sh",
		},
	}

	if err := c.NewRecords("peanuts.org", recs); err != nil {
		t.Error(err)
	}
	if recs == nil {
		t.Error("no records returned, expected 1")
	}
	r := *recs
	if len(r) == 0 {
		t.Error("no records returned, expected 1")
	}
	logger.Debug(r[0].ID)
}

func TestDeleteRecordsWithInvalidID(t *testing.T) {
	setTestKey(t)
	if err := c.DeleteRecords("peanuts.org", &SRVRecords{{ID: "testid"}}); err == nil {
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
	if err := c.GetRecords("peanuts.org", r); err != nil {
		t.Errorf("unexpected error returned: %v", err)
	}
	for _, rec := range *r {
		logger.Debug(rec.ID)
		logger.Debug(rec.Name)
	}

	if err := c.DeleteRecords("peanuts.org", &SRVRecords{{ID: "dGVzdGluZy5wZWFudXRzLm9yZy9DTkFNRXw4NjQ"}}); err != nil {
		t.Errorf("unexpected error returned: %v", err)
	}
}

func TestCreateARecord(t *testing.T) {
	setTestKey(t)
	if err := c.NewRecords("peanuts.org", &ARecords{{Name: "testing-a-record", Target: "179.61.181.1"}}); err != nil {
		t.Errorf("unexpected error returned: %v", err)
	}
}

func TestCreateCNAMERecord(t *testing.T) {
	setTestKey(t)
	if err := c.NewRecords("peanuts.org", &CNAMERecords{{Name: "testing-cname-record", Target: "179.61.181.1"}}); err != nil {
		t.Errorf("unexpected error returned: %v", err)
	}
}

func TestGetCNAMERecords(t *testing.T) {
	setTestKey(t)
	cnames := &CNAMERecords{}
	if err := c.GetRecords("peanuts.org", cnames); err != nil {
		t.Error(err)
	}
	logger.Debug(cnames)
}
