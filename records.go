package netacl

import (
	"encoding/json"
	"errors"
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
	ID  string  `json:"id,omitempty"`
	Obj *record `json:"obj,omitempty"`
}

type record struct {
	Name string `json:"name"`
	Data struct {
		SRV *SRVRecord `json:"SRV,omitempty"`
	} `json:"data"`
}

type SRVRecord struct {
	ID   string `json:"-"`
	Name string `json:"-"`

	Target   string `json:"target"`
	Port     int16  `json:"port"`
	Priority int16  `json:"priority"`
	Weight   int16  `json:"weight"`
}

type SRVRecords []*SRVRecord

func (r *SRVRecords) add(domain string, c *APICLient) error {
	if r == nil {
		return errors.New("records cannot be nil")
	}
	payload := []map[string]addPayload{}
	for _, rec := range *r {
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

	raw, err := c.Request(fmt.Sprintf("/dns/%v", domain), http.MethodPatch, Application_json, payload)
	if err != nil {
		return err
	}

	added := SRVRecords{}
	result := []map[string]addPayload{}
	if err := json.Unmarshal(raw, &result); err != nil {
		return err
	}
	for _, res := range result {
		if data, ok := res["Added"]; ok {
			rec := data.Obj.Data.SRV

			// set the two unset fields
			rec.ID = data.ID
			rec.Name = data.Obj.Name

			added = append(added, rec)
		}
	}

	*r = added
	return nil
}
