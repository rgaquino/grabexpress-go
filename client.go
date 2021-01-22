package grabexpress

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Client may be used to make requests to the GrabExpress APIs
type Client struct {
	httpClient *http.Client
	apiKey     string
	secret     string
	baseURL    string
	tokenURL   string
	oauth      clientcredentials.Config
}

// DTO ...
type DTO interface {
	SetRequestID(id string)
}

// ClientOption is the type of constructor options for NewClient(...).
type ClientOption func(*Client) error

// NewClient constructs a new Client which can make requests to the GrabExpress APIs.
func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	if strings.TrimSpace(c.apiKey) == "" || strings.TrimSpace(c.secret) == "" {
		return nil, ErrCredentialsMissing
	}
	if strings.TrimSpace(c.baseURL) == "" {
		return nil, ErrBaseURLMissing
	}
	if strings.TrimSpace(c.baseURL) == "" {
		return nil, ErrTokenURLMissing
	}
	c.oauth = clientcredentials.Config{
		ClientID:     c.apiKey,
		ClientSecret: c.secret,
		TokenURL:     c.tokenURL,
		Scopes:       []string{"grab_express.partner_deliveries"},
	}
	return c, nil
}

// WithHTTPClient configures a GrabExpress API client with a http.Client to make requests over.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		if c.Transport == nil {
			c.Transport = http.DefaultTransport
		}
		client.httpClient = c
		return nil
	}
}

// WithAPIKey configures a GrabExpress API client with an API Key
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) error {
		c.apiKey = apiKey
		return nil
	}
}

// WithSecret configures a GrabExpress API client with a secret
func WithSecret(secret string) ClientOption {
	return func(c *Client) error {
		c.secret = secret
		return nil
	}
}

// WithBaseURL configures a GrabExpress API client with a custom base url
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// WithTokenURL configures a GrabExpress API client with an auth token url
func WithTokenURL(tokenURL string) ClientOption {
	return func(c *Client) error {
		c.tokenURL = tokenURL
		return nil
	}
}

func (c *Client) get(ctx context.Context, path string, apiReq interface{}, apiResp DTO) error {
	req, err := c.createRequest(ctx, http.MethodGet, path, apiReq)
	if err != nil {
		return wrapError(err)
	}
	return c.do(ctx, req, apiResp)
}

func (c *Client) post(ctx context.Context, path string, apiReq interface{}, apiResp DTO) error {
	req, err := c.createRequest(ctx, http.MethodPost, path, apiReq)
	if err != nil {
		return wrapError(err)
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(ctx, req, apiResp)
}

func (c *Client) put(ctx context.Context, path string, apiReq interface{}, apiResp DTO) error {
	req, err := c.createRequest(ctx, http.MethodPut, path, apiReq)
	if err != nil {
		return wrapError(err)
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(ctx, req, apiResp)
}

func (c *Client) delete(ctx context.Context, path string, apiReq interface{}, apiResp DTO) error {
	req, err := c.createRequest(ctx, http.MethodDelete, path, apiReq)
	if err != nil {
		return wrapError(err)
	}
	return c.do(ctx, req, apiResp)
}

func (c *Client) createRequest(ctx context.Context, method, path string, apiReq interface{}) (*http.Request, error) {
	body, err := marshalRequest(apiReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}
	token, err := c.generateAuth(ctx)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, apiResp DTO) error {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return wrapError(err)
	}
	defer resp.Body.Close()
	return decodeResponse(resp, apiResp)
}

func (c *Client) generateAuth(ctx context.Context) (*oauth2.Token, error) {
	token, err := c.oauth.Token(ctx)
	if err != nil {
		return nil, ErrAuthenticationError
	}
	return token, nil
}

func marshalRequest(apiReq interface{}) (io.Reader, error) {
	if apiReq == nil {
		return nil, nil
	}
	body, err := json.Marshal(apiReq)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

func decodeResponse(resp *http.Response, apiResp DTO) error {
	requestID := resp.Header.Get("X-Grabkit-Grab-Requestid")
	apiResp.SetRequestID(requestID)

	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(apiResp); err != nil {
			return wrapError(err)
		}
		return nil
	case http.StatusNoContent:
		return nil
	default:
		var msg string
		if resp.ContentLength != 0 {
			bb, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return wrapError(err)
			}
			msg = string(bb)
		}
		return &Error{
			Status:    resp.StatusCode,
			Message:   msg,
			RequestID: requestID,
		}
	}
}
