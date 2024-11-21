package netacl

import (
	"encoding/json"
	"errors"
)

var (
	ErrRecordExists error = errors.New("record with the same data already exists")
)

type apierror struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

func extract(raw []byte) error {
	e := &apierror{}
	if err := json.Unmarshal(raw, e); err != nil {
		return nil
	}
	switch e.Code {
	case 568543:
		return ErrRecordExists
	default:
		if e.Desc == "" {
			return nil
		}
		return errors.New(e.Desc)
	}
}
