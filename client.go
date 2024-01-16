package pocket

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const defaultTimeout = 5 * time.Second
const host string = "https://getpocket.com/v3"

type Client struct {
	hc   *http.Client
	Auth Auth
}

func NewClient(params *AuthParams) (*Client, error) {
	if params.ConsumerKey == "" {
		return nil, errors.New("consumer key is empty")
	}

	c := &Client{
		hc: &http.Client{
			Timeout: defaultTimeout,
		},
		Auth: Auth{
			AuthParams: params,
		},
	}

	c.Auth.client = c
	return c, nil
}

func (c *Client) do(ctx context.Context, method string, endpoint string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, host+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create new request")
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF8")

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to send http request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("API Error: (%s) %s", resp.Header.Get("X-Error-Code"), resp.Header.Get("X-Error"))
		return nil, errors.New(err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to read request body")
	}

	return body, nil
}

func (c *Client) post(ctx context.Context, endpoint string, body []byte) ([]byte, error) {
	return c.do(ctx, http.MethodPost, endpoint, body)
}

func (c *Client) get(ctx context.Context, endpoint string, body []byte) ([]byte, error) {
	return c.do(ctx, http.MethodGet, endpoint, body)
}
