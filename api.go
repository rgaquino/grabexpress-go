package grabexpress

import (
	"context"
	"fmt"
)

// CreateQuotes requests for delivery service quotes. When packages details aren't provided,
// a single cheapest category package is assumed. Immediate dispatching is assumed.
// An array of delivery services with their respective quote is returned.
func (c *Client) CreateQuotes(ctx context.Context, req *CreateQuotesRequest) (*CreateQuotesResponse, error) {
	path := "/v1/deliveries/quotes"
	resp := &CreateQuotesResponse{}
	if err := c.post(ctx, path, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateDelivery ...
func (c *Client) CreateDelivery(ctx context.Context, req *CreateDeliveryRequest) (*CreateDeliveryResponse, error) {
	path := "/v1/deliveries"
	resp := &CreateDeliveryResponse{}
	if err := c.post(ctx, path, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDelivery ...
func (c *Client) GetDelivery(ctx context.Context, deliveryID string) (*GetDeliveryResponse, error) {
	path := fmt.Sprintf("/v1/deliveries/%s", deliveryID)
	resp := &GetDeliveryResponse{}
	if err := c.get(ctx, path, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// CancelDelivery ...
func (c *Client) CancelDelivery(ctx context.Context, deliveryID string) error {
	path := fmt.Sprintf("/v1/deliveries/%s", deliveryID)
	resp := &GetDeliveryResponse{}
	if err := c.delete(ctx, path, nil, resp); err != nil {
		return err
	}
	return nil
}
