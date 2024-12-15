package netacl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"

	"github.com/roolps/logging"
)

var c = &APICLient{}

var logger = &logging.Profile{
	Color:  color.RGB(255, 128, 0),
	Prefix: "NetACL",
}

const (
	Application_json = "application/json"
	Text_plain       = "text/plain"
)

type APICLient struct{ Secret string }

func EnableDebug() {
	logger.EnableDebug()
}

func NewClient(apikey string) (*APICLient, error) {
	if apikey == "" {
		return nil, errors.New("api key cannot be empty")
	}
	return &APICLient{Secret: apikey}, nil
}

// returns raw request body and any errors
func (c *APICLient) Request(endpoint, method, contentType string, body any) ([]byte, error) {
	// initialise variables
	var (
		req *http.Request
		err error
		raw []byte

		client = &http.Client{}
	)
	if body != nil {
		raw, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %v", err)
		}
	}

	// create http request
	logger.Debugf("%v:%v", method, endpoint)
	req, err = http.NewRequest(method, fmt.Sprintf("https://netacl.com/api/%v", endpoint), bytes.NewBuffer(raw))
	if err != nil {
		return nil, err
	}

	// add relevant headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Secret))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// extract all data from body
	raw, err = io.ReadAll(res.Body)
	if err != nil {
		// eof is for empty body
		if err != io.EOF {
			return nil, err
		}
	}

	if res.StatusCode >= 400 {
		if err = extract(raw); err == nil {
			return nil, fmt.Errorf("[%v] %v", res.StatusCode, string(raw))
		} else {
			return nil, err
		}
	}
	return raw, nil
}
