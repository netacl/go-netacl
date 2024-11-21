package netacl

import (
	"encoding/json"
	"net/http"
)

type Domains struct {
	Owned   []string `json:"owned"`
	Managed []string `json:"managed"`
}

func (c *APICLient) GetDomains() (*Domains, error) {
	raw, err := c.Request("/dns", http.MethodGet, Application_json, nil)
	if err != nil {
		return nil, err
	}
	data := &Domains{}
	json.Unmarshal(raw, data)

	// add back error handling later when object received as correct type
	// if err := json.Unmarshal(raw, data); err != nil {
	// 	return nil, err
	// }

	return data, nil
}
