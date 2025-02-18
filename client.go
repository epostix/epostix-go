package epostix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.epostix.com"
)

type client struct {
	apiKey  string
	options Options
}

func (c *client) SendEmail(ctx context.Context, domainName string, emailCreate *EmailCreate) (*Email, error) {
	b, err := json.Marshal(emailCreate)
	if err != nil {
		return nil, err
	}

	finalURL := fmt.Sprintf("%s/%s/emails", baseURL, domainName)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, finalURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.apiKey)

	opts := c.options
	resp, err := opts.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to send email: %s (status: %d)", string(body), resp.StatusCode)
	}

	var email *Email
	if err := json.NewDecoder(resp.Body).Decode(&email); err != nil {
		return nil, err
	}

	return email, nil
}

func New(apiKey string, opts ...Option) Client {
	options := Options{
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &client{
		apiKey:  apiKey,
		options: options,
	}
}
