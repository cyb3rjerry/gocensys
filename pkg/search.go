package censys

import (
	"net/url"
	"strconv"
)

type HostSearchQuery struct {
	Query       string   `json:"query"`
	VirtualHost string   `json:"virtual_hosts,omitempty"`
	Sort        string   `json:"sort"`
	Cursor      string   `json:"cursor"`
	Fields      []string `json:"fields"`
	PerPage     int      `json:"per_page"`
}

type SearchResponse struct {
	Query string `json:"query"`
	Links struct {
		Next string `json:"next"`
		Prev string `json:"prev"`
	} `json:"links"`
	Hits     []SearchHit `json:"hits"`
	Total    float64     `json:"total"`
	Duration int         `json:"duration_ms"`
}

type SearchHit struct {
	Services []Services `json:"services,omitempty"`
	Dns      struct {
		ReverseDNS struct {
			Names []string `json:"names,omitempty"`
		} `json:"reverse_dns,omitempty"`
	} `json:"dns,omitempty"`
	ParsedCertificate struct {
		ValidityPeriod struct {
			NotAfter  string `json:"not_after,omitempty"`
			NotBefore string `json:"not_before,omitempty"`
		} `json:"validity_period,omitempty"`
		SubjectDN string `json:"subject_dn,omitempty"`
		IssuerDN  string `json:"issuer_dn,omitempty"`
	} `json:"parsed,omitempty"`
	LastUpdated      string           `json:"last_updated,omitempty"`
	IP               string           `json:"ip,omitempty"`
	SHA256           string           `json:"fingerprint_sha256,omitempty"`
	OS               HostOS           `json:"operating_system,omitempty"`
	AutonomousSystem AutonomousSystem `json:"autonomous_system,omitempty"`
	Location         Location         `json:"location,omitempty"`
	Names            []string         `json:"names,omitempty"`
}

type Services struct {
	Protocol            string `json:"transport_protocol,omitempty"`
	ServiceName         string `json:"service_name,omitempty"`
	ExtendedServiceName string `json:"extended_service_name,omitempty"`
	Certificate         string `json:"certificate,omitempty"`
	Port                int    `json:"port,omitempty"`
}

type HostOS struct {
	CPE       string `json:"cpe,omitempty"`
	Partition string `json:"part,omitempty"`
	Vendor    string `json:"vendor,omitempty"`
	Source    string `json:"source,omitempty"`
	Product   string `json:"product,omitempty"`
}

type Location struct {
	PostalCode  string `json:"postal_code,omitempty"`
	Province    string `json:"province,omitempty"`
	Country     string `json:"country,omitempty"`
	City        string `json:"city,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	Continent   string `json:"continent,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Coordinates struct {
		Longitude float64 `json:"longitude,omitempty"`
		Latitude  float64 `json:"latitude,omitempty"`
	} `json:"coordinates,omitempty"`
}

type AutonomousSystem struct {
	CountryCode string `json:"country_code,omitempty"`
	BGPPrefix   string `json:"bgp_prefix,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ASN         int    `json:"asn,omitempty"`
}

type CertSearchQuery struct {
	Query   string   `json:"query,omitempty"`
	Sort    []string `json:"sort,omitempty"`
	Cursor  string   `json:"cursor,omitempty"`
	PerPage int      `json:"per_page,omitempty"`
}

func (c *Client) NewHostSearchQuery() *HostSearchQuery {
	return &HostSearchQuery{
		Query:       "",
		VirtualHost: "INCLUDE",
		Sort:        "RELEVANCE",
		Cursor:      "",
		Fields:      []string{},
		PerPage:     50,
	}
}

func (c *Client) NewCertSearchQuery() *CertSearchQuery {
	return &CertSearchQuery{
		Query:   "",
		Sort:    []string{"parsed.issuer.organization", "parsed.subject.country"},
		Cursor:  "",
		PerPage: 50,
	}
}

func (c *Client) DoHostSearch(query *HostSearchQuery) (*SearchResponse, error) {
	resp := baseResponse{
		Results: &SearchResponse{},
	}

	params := url.Values{
		"q":             {query.Query},
		"per_page":      {strconv.Itoa(query.PerPage)},
		"fields":        query.Fields,
		"virtual_hosts": {query.VirtualHost},
		"sort":          {query.Sort},
		"cursor":        {query.Cursor},
	}

	req, err := c.NewRequest("/hosts/search?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}

	return resp.Results.(*SearchResponse), nil
}

func (c *Client) DoCertSearch(query *CertSearchQuery) (*SearchResponse, error) {
	resp := baseResponse{
		Results: &SearchResponse{},
	}
	params := url.Values{
		"q":        {query.Query},
		"per_page": {strconv.Itoa(query.PerPage)},
		"sort":     query.Sort,
		"cursor":   {query.Cursor},
	}

	req, err := c.NewRequest("/certificates/search?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}
	return resp.Results.(*SearchResponse), nil
}
