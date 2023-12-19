package censys

import (
	"io"
	"net/http"
	"net/url"
)

const v1PathURL = "https://search.censys.io/api/v1"

type V1AccountInfo struct {
	Email      string `json:"email"`
	Login      string `json:"login"`
	FirstLogin string `json:"first_login"`
	LastLogin  string `json:"last_login"`
	Quota      struct {
		ResetsAt  string `json:"resets_at"`
		Used      int    `json:"used"`
		Allowance int    `json:"allowance"`
	} `json:"quota"`
}

func (c *Client) NewV1Request(path string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(v1PathURL + path)
	if err != nil {
		return nil, err
	}

	return c.newRequest(u, body)
}

func (c *Client) GetAccountInfo() (*V1AccountInfo, error) {
	resp := V1AccountInfo{}
	req, err := c.NewV1Request("/account", nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
