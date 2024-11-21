package netacl

import "testing"

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
			Name:     "_minecraft._tcp.extra.mcsh.io",
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
