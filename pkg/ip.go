package censys

import "net/url"

type IP struct {
	IP                        string `json:"ip"`
	AutonomousSystemUpdatedAt string `json:"autonomous_system_updated_at"`
	LocationUpdatedAt         string `json:"location_updated_at"`
	LastUpdated               string `json:"last_updated"`
	// TODO: Services is a list of any, but I'm not sure who to define it's type
	// as it can be a bunch of different things
	Services         []any      `json:"services"`
	Location         IPLocation `json:"location"`
	AutonomousSystem IPASN      `json:"autonomous_system"`
	OS               IPOS       `json:"operation_system"`
	DNS              IPDNS      `json:"dns"`
}

type IPLocation struct {
	Continent   string `json:"continent"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	Timezone    string `json:"timezone"`
	Province    string `json:"province"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
}

type IPASN struct {
	Description string `json:"description"`
	BgpPrefix   string `json:"bgp_prefix"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	ASN         int    `json:"asn"`
}

type IPOS struct {
	URI     string `json:"uniform_resource_identifier"`
	Part    string `json:"part"`
	Product string `json:"product"`
	Source  string `json:"source"`
}

type IPDNS struct {
	Names   []string `json:"names"`
	Records map[string]struct {
		RecordType string `json:"record_type"`
		ResolvedAt string `json:"resolved_at"`
	} `json:"records"`
	ReverseDNS struct {
		ResolvedAt string   `json:"resolved_at"`
		Names      []string `json:"names"`
	} `json:"reverse_dns"`
}

type IPQuery struct {
	IP     string
	AtTime string
}

func (c *Client) NewIPQuery(IP string) *IPQuery {
	return &IPQuery{
		IP:     IP,
		AtTime: "",
	}
}

func (c *Client) DoIPQuery(query *IPQuery) (*IP, error) {
	resp := baseResponse{
		Results: &IP{},
	}

	params := url.Values{
		"at_time": {query.AtTime},
	}

	req, err := c.NewRequest("/hosts/"+query.IP+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}

	return resp.Results.(*IP), nil
}
