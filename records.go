package netacl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Records interface {
	get(string, *APICLient) error
	add(string, *APICLient) error
	remove(string, *APICLient) error
}

// create new dns record
func (c *APICLient) NewRecords(domain string, r Records) error {
	return r.add(domain, c)
}

// create new dns record
func (c *APICLient) DeleteRecords(domain string, r Records) error {
	return r.remove(domain, c)
}

func (c *APICLient) GetRecords(domain string, r Records) error {
	return r.get(domain, c)
}

type payload struct {
	ID  string  `json:"id,omitempty"`
	Obj *record `json:"obj,omitempty"`
}

type record struct {
	Name string `json:"name"`
	Data struct {
		SRV *SRVRecord `json:"SRV,omitempty"`
	} `json:"data"`
}

type SRVRecords []*SRVRecord

type SRVRecord struct {
	ID   string `json:"-"`
	Name string `json:"-"`

	Target   string `json:"target"`
	Port     int16  `json:"port"`
	Priority int16  `json:"priority"`
	Weight   int16  `json:"weight"`
}

func (r *SRVRecords) get(domain string, c *APICLient) error {
	raw, err := c.Request(fmt.Sprintf("/dns/%v", domain), http.MethodGet, Application_json, nil)

	result := map[string]record{}
	if err := json.Unmarshal(raw, &result); err != nil {
		return err
	}

	records := SRVRecords{}
	for id, res := range result {
		// need to find a nice way to unmarshal this payload
		// this is nil because there is no "SRV" key, the data is just the object...
		// rec := res.Data.SRV
		// rec.ID = id
		// rec.Name = res.Name

		records = append(records, &SRVRecord{ID: id, Name: res.Name})
	}
	*r = records
	return err
}

func (r *SRVRecords) add(domain string, c *APICLient) error {
	if r == nil {
		return errors.New("records cannot be nil")
	}
	pl := []map[string]payload{}
	for _, rec := range *r {
		pl = append(pl, map[string]payload{
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

	raw, err := c.Request(fmt.Sprintf("/dns/%v", domain), http.MethodPatch, Application_json, pl)
	if err != nil {
		return err
	}

	result := []map[string]payload{}
	if err := json.Unmarshal(raw, &result); err != nil {
		return err
	}

	records := SRVRecords{}
	for _, added := range result {
		if res, ok := added["Added"]; ok {
			rec := res.Obj.Data.SRV

			rec.ID = res.ID
			rec.Name = res.Obj.Name

			records = append(records, rec)
		}
	}
	*r = records
	return nil
}

func (r *SRVRecords) remove(domain string, c *APICLient) error {
	if r == nil {
		return errors.New("records cannot be nil")
	}
	pl := []map[string]payload{}
	for _, rec := range *r {
		pl = append(pl, map[string]payload{
			"Remove": {
				ID: rec.ID,
			},
		})
	}

	_, err := c.Request(fmt.Sprintf("/dns/%v", domain), http.MethodPatch, Application_json, pl)
	return err
}
