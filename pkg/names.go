package censys

import (
	"net/url"
	"strconv"
)

type IPNames struct {
	IP    string   `json:"ip"`
	Names []string `json:"names"`
}

type IPNamesQuery struct {
	IP      string
	Cursor  string
	PerPage int
}

func (c *Client) NewIPNamesQuery() *IPNamesQuery {
	return &IPNamesQuery{
		IP: "",
	}
}

func (c *Client) DoIPNameQuery(q *IPNamesQuery) (*IPNames, error) {
	resp := baseResponse{
		Results: &IPNames{},
	}

	params := url.Values{
		"per_page": {strconv.Itoa(q.PerPage)},
		"cursor":   {q.Cursor},
	}

	req, err := c.NewRequest("/hosts/"+q.IP+"/names?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}

	return resp.Results.(*IPNames), nil
}
