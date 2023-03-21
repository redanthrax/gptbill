package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

const (
    CONGRESSURL = "https://api.congress.gov/v3"
)

type Client struct {
    BaseURL string
    ApiKey string
    HTTPClient *http.Client
}

type service struct {
    Client
}

func CongressApi() (*Client, error) {
    apiKey := os.Getenv("CONGRESSAPI")
    if apiKey == "" {
        return nil, errors.New("CONGRESSAPI api key missing from env vars")
    }

    return &Client{
        BaseURL: CONGRESSURL,
        ApiKey: apiKey,
        HTTPClient: &http.Client{
            Timeout: time.Minute,
        },
    }, nil
}

func (c *Client) CongressRequest(req *http.Request, v interface{}) error {
    q := req.URL.Query()
	q.Add("api_key", c.ApiKey)
	req.URL.RawQuery = q.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
