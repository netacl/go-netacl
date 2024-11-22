package netacl

import (
	"encoding/json"
	"errors"
)

var (
	ErrRecordExists            error = errors.New("same record data already exists")
	ErrRecordDoesntExistInZone error = errors.New("the record doesn't belong to the zone")
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
		switch e.Desc {
		case "The record does not belong to the zone":
			return ErrRecordDoesntExistInZone
		case "Same Record Data already exists":
			return ErrRecordExists
		}
		return ErrRecordExists
	default:
		if e.Desc == "" {
			return nil
		}
		return errors.New(e.Desc)
	}
}
