package epostix

import (
	"context"
	"net/http"
)

type Client interface {
	SendEmail(ctx context.Context, domainName string, emailCreate *EmailCreate) (*Email, error)
}

type Options struct {
	httpClient *http.Client
}

type Option func(*Options)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(o *Options) {
		o.httpClient = httpClient
	}
}
