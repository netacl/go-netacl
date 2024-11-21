package netacl

import (
	"log"
	"testing"
)

// test getting a list of domains
func TestGetDomains(t *testing.T) {
	setTestKey(t)
	if d, err := c.GetDomains(); err != nil {
		t.Error(err)
	} else {
		log.Println(d)
	}
}
