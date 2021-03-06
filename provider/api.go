package provider

import (
	"net"
)

// API is ...
type API struct {
	factoryServer Daemon
}

// NewAPI is ...
func NewAPI(opts ...APIOption) *API {
	provider := &API{}

	for _, opt := range opts {
		opt(provider)
	}

	return provider
}

// APIOption is ...
type APIOption func(*API)

// WithFactoryServer is ...
func WithFactoryServer(i Daemon) APIOption {
	return func(s *API) {
		if i != nil {
			s.factoryServer = i
		}
	}
}

// Daemon is ...
type Daemon interface {
	Start() net.Listener
}

// Activate is ...
func (s *API) Activate() {
	s.factoryServer.Start()
}
