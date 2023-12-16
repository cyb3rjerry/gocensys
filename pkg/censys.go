package censys

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Endpoint string

const (
	// Censys v2 API base URL
	BaseURL = "https://search.censys.io/api/v2"
)

type Client struct {
	ApiID     string
	ApiSecret string
	Client    *http.Client
	BaseURL   string
}

// NewClient creates a new Censys API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(apiID, apiSecret string) *Client {
	return &Client{
		ApiID:     apiID,
		ApiSecret: apiSecret,
		Client:    http.DefaultClient,
		BaseURL:   BaseURL,
	}
}

// BaseResponse is the base response from Censys API.
// This format is used for all responses.
type baseResponse struct {
	Results any    `json:"result"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

// NewRequest creates an API request. A relative URL can be provided in path
func (c *Client) NewRequest(path string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}

	return c.newRequest(u, body)
}

// newRequest creates an API request. A relative URL can be provided in urlStr,
func (c *Client) newRequest(url *url.URL, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.ApiID, c.ApiSecret)
	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

// Do sends an API request and returns an error if one rises.
func (c *Client) Do(ctx context.Context, req *http.Request, dest interface{}) error {
	req = req.WithContext(ctx)
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status error: %d", resp.StatusCode)
	}

	if dest == nil {
		return errors.New("can't unmarshal response to nil struct")
	}

	return c.parseResponse(resp.Body, dest)
}

// parseResponse parses the response body into the destination interface.
func (c *Client) parseResponse(r io.Reader, dest interface{}) error {
	var err error
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r)
	if err != nil {
		return err
	}

	s := buf.String()

	if w, ok := dest.(io.Writer); ok {
		_, err = io.Copy(w, buf)
	} else {
		err = json.Unmarshal([]byte(s), dest)
	}
	return err
}
