package censys

import (
	"net/url"
)

type CertHosts struct {
	Fingerprints string `json:"fingerprint"`
	Hosts        []struct {
		IP              string `json:"ip"`
		Name            string `json:"name"`
		ObservedAt      string `json:"observed_at"`
		FirstObservedAt string `json:"first_observed_at"`
	} `json:"hosts"`
}

type CertHostsQuery struct {
	Fingerprint string
	Cursor      string
}

func (c *Client) NewCertHostsQuery(fingerprint string) *CertHostsQuery {
	return &CertHostsQuery{Fingerprint: fingerprint}
}

func (c *Client) DoCertHostsQuery(query *CertHostsQuery) (*CertHosts, error) {
	resp := baseResponse{
		Results: &CertHosts{},
	}
	params := url.Values{
		"cursor": {query.Cursor},
	}

	req, err := c.NewRequest("/certificates/"+query.Fingerprint+"/hosts?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}
	return resp.Results.(*CertHosts), nil
}
