package grabexpress

import (
	"context"
)

// CreateQuotes requests for delivery service quotes. When packages details aren't provided,
// a single cheapest category package is assumed. Immediate dispatching is assumed.
// An array of delivery services with their respective quote is returned.
func (c *Client) CreateQuotes(ctx context.Context, req *CreateQuotesRequest) (*CreateQuotesResponse, error) {
	path := "/v2/quotations"
	resp := &CreateQuotesResponse{}
	if err := c.post(ctx, path, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}