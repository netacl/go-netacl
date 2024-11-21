package netacl

import (
	"fmt"
	"net/http"
)

type Records interface {
	add(string, *APICLient) error
}

// create new dns record
func (c *APICLient) NewRecords(domain string, r Records) error {
	return r.add(domain, c)
}

type addPayload struct {
	Obj *record `json:"obj,omitempty"`
}

type record struct {
	Name string `json:"name"`
	Data struct {
		SRV *SRVRecord `json:"SRV,omitempty"`
	} `json:"data"`
}

type SRVRecord struct {
	Name string `json:"-"`

	Target   string `json:"target"`
	Port     int16  `json:"port"`
	Priority int16  `json:"priority"`
	Weight   int16  `json:"weight"`
}

type SRVRecords []*SRVRecord

func (r SRVRecords) add(domain string, c *APICLient) error {
	payload := []map[string]addPayload{}
	for _, rec := range r {
		payload = append(payload, map[string]addPayload{
			"Add": {
				Obj: &record{
					Name: rec.Name,
					Data: struct {
						SRV *SRVRecord "json:\"SRV,omitempty\""
					}{
						SRV: rec,
					},
				},
			},
		})
	}
	_, err := c.Request(fmt.Sprintf("/dns/%v", domain), http.MethodPatch, Application_json, payload)
	return err
}
