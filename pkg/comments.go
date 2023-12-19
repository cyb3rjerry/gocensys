package censys

type IPComment struct {
	IP       string `json:"ip"`
	Comments []struct {
		Id        string `json:"id"`
		IP        string `json:"ip"`
		AuthorId  string `json:"author_id"`
		Contents  string `json:"contents"`
		CreatedAt string `json:"created_at"`
	} `json:"comments"`
}

type IPCommentQuery struct {
	IP string
}

func (c *Client) NewIPCommentQuery() *IPCommentQuery {
	return &IPCommentQuery{
		IP: "",
	}
}

func (c *Client) DoIPCommentQuery(q *IPCommentQuery) (*IPComment, error) {
	resp := baseResponse{
		Results: &IPComment{},
	}
	req, err := c.NewRequest("/hosts/"+q.IP+"/comments", nil)
	if err != nil {
		return nil, err
	}
	if err := c.Do(req.Context(), req, &resp); err != nil {
		return nil, err
	}
	return resp.Results.(*IPComment), nil
}
