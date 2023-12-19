package censys

import "net/url"

type IPDiff struct {
	IPA struct {
		IP          string `json:"ip"`
		LastUpdated string `json:"last_updated"`
	} `json:"a"`
	IPB struct {
		IP          string `json:"ip"`
		LastUpdated string `json:"last_updated"`
	} `json:"b"`
	Patch []Patch `json:"patch"`
}

type Patch struct {
	Value any    `json:"value"`
	OP    string `json:"op"`
	Path  string `json:"path"`
}

type DiffQuery struct {
	IPA     string
	IPB     string
	AtTimeA string
	AtTimeB string
}

func (c *Client) NewIPDiffQuery() *DiffQuery {
	return &DiffQuery{
		IPA:     "",
		IPB:     "",
		AtTimeA: "",
		AtTimeB: "",
	}
}

func (c *Client) DoIPDiffQuery(q *DiffQuery) (*IPDiff, error) {
	resp := baseResponse{
		Results: &IPDiff{},
	}
	params := url.Values{
		"ip_b":      {q.IPB},
		"at_time_a": {q.AtTimeA},
		"at_time_b": {q.AtTimeB},
	}
	req, err := c.NewRequest("/hosts/"+q.IPA+"/diff?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}
	return resp.Results.(*IPDiff), nil
}
