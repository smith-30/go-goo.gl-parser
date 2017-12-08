package parser

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	baseHost       = "goo.gl"
	requestBaseUrl = "https://www.googleapis.com/urlshortener/v1/url"
)

type (
	Parser struct {
		APIKey string
		Fetcher
	}
)

func NewParser(apiKey string) Parser {
	return Parser{
		APIKey:  apiKey,
		Fetcher: Client{},
	}
}

func (p Parser) DecodeURL(target string) (string, error) {
	u, err := url.Parse(target)
	if err != nil {
		return "", err
	}

	if u.Host != baseHost {
		return "", errors.New(fmt.Sprintf("url format is invalid: %s", u.Host))
	}

	req, _ := url.Parse(requestBaseUrl)

	q := req.Query()
	q.Set("key", p.APIKey)
	q.Set("shortUrl", target)
	req.RawQuery = q.Encode()

	res, err := p.fetch(req.String())
	if err != nil {
		return "", err
	}

	return res.LongURL, nil
}
