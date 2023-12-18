package censys

import (
	"fmt"
	"net/url"
	"strconv"
)

type Aggregate struct {
	Query   string `json:"query"`
	Field   string `json:"field"`
	Buckets []struct {
		Key   string `json:"key"`
		Count int    `json:"count"`
	} `json:"buckets"`
	TotalResults       float64 `json:"total"`
	Duration           int     `json:"duration"`
	TotalOmitted       int     `json:"total_omitted"`
	TotalNested        int     `json:"total_nested"`
	PotentialDeviation int     `json:"potential_deviation"`
}

type AggregateQuery struct {
	Query       string
	Field       string
	VirtualHost string
	NumBuckets  int
}

func (c *Client) NewAggregateQuery() *AggregateQuery {
	return &AggregateQuery{
		Query:       "",
		Field:       "",
		VirtualHost: "INCLUDE",
		NumBuckets:  50,
	}
}

func (c *Client) DoHostAggregateQuery(query *AggregateQuery) (*Aggregate, error) {
	resp := baseResponse{
		Results: &Aggregate{},
	}

	if query.Field == "" {
		query.Field = "services.port"
	}

	params := url.Values{
		"q":             {query.Query},
		"field":         {query.Field},
		"num_buckets":   {strconv.Itoa(query.NumBuckets)},
		"virtual_hosts": {query.VirtualHost},
	}

	req, err := c.NewRequest("/hosts/aggregate?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}
	return resp.Results.(*Aggregate), nil
}

func (c *Client) DoCertificateAggregateQuery(query *AggregateQuery) (*Aggregate, error) {
	resp := baseResponse{
		Results: &Aggregate{},
	}

	if query.Field == "" {
		query.Field = "parsed.issuer.organization"
	}

	params := url.Values{
		"q":           {query.Query},
		"field":       {query.Field},
		"num_buckets": {strconv.Itoa(query.NumBuckets)},
	}

	req, err := c.NewRequest("/certificates/aggregate?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", req)
	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}
	return resp.Results.(*Aggregate), nil
}
