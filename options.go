package postmark

import "net/http"

// Option to set an optional client value.
type Option func(o *Options)

// Options for our client.
type Options struct {
	Client     *http.Client
	BackendURL string
	UserAgent  string
}

// NewOptions returns a new Options with defaults and supplied overrides.
func NewOptions(opts ...Option) *Options {
	out := Options{
		Client:     http.DefaultClient,
		BackendURL: backendURL,
		UserAgent:  userAgent,
	}

	for _, o := range opts {
		o(&out)
	}

	return &out
}

// WithClient ...
func WithClient(v *http.Client) Option {
	return func(o *Options) {
		o.Client = v
	}
}

// WithBackendURL ...
func WithBackendURL(v string) Option {
	return func(o *Options) {
		o.BackendURL = v
	}
}

// WithUserAgent ...
func WithUserAgent(v string) Option {
	return func(o *Options) {
		o.UserAgent = v
	}
}
