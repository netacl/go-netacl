package netacl

import (
	"encoding/json"
	"net/http"
)

type Domains struct {
	Owned []string `json:"owned"`

	// map keys are type 'username (email)'
	Managed map[string][]string `json:"managed"`
}

func (c *APICLient) GetDomains() (*Domains, error) {
	raw, err := c.Request("/dns", http.MethodGet, Application_json, nil)
	if err != nil {
		return nil, err
	}
	data := &Domains{}
	if err := json.Unmarshal(raw, data); err != nil {
		return nil, err
	}
	return data, nil
}
