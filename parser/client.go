package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	Fetcher interface {
		fetch(url string) (*Response, error)
	}

	Client struct {
	}
)

func (c Client) fetch(url string) (*Response, error) {
	r := &Response{}

	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return r, errors.New(fmt.Sprintf("failed HTTPClient.Do: %v", err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r, errors.New(fmt.Sprintf("resp body read err: %v", err))
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return r, errors.New(fmt.Sprintf("failed parse json: %v", err))
	}

	return r, nil
}
