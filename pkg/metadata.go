package censys

type Metadata struct {
	Services []string `json:"services"`
}

// The host metadata endpoint returns a list of services Censys scans for.
// These are the values that can be given as values for the services.service_name
// field in search queries.
func (c *Client) GetHostMetadata() (*Metadata, error) {
	resp := baseResponse{
		Results: &Metadata{},
	}

	req, err := c.NewRequest("/metadata/hosts", nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}

	return resp.Results.(*Metadata), nil
}
