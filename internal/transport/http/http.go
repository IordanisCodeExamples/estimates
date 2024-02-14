// Package transporthttp provides the HTTP transport layer for the estimates service.
// It includes the setup and management of the HTTP server, handling of routes and requests.
package transporthttp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	merror "github.com/junkd0g/neji"
	"github.com/rs/cors"

	"github.com/junkd0g/estimates/internal/transport"
)

// HTTPServer wraps all the components necessary for running an HTTP server.
// It includes logging, business logic services, and server configuration.
type HTTPServer struct {
	Logger       transport.Logger
	Service      transport.Service
	ServerConfig *ServerConfig
}

// New initializes a new instance of HTTPServer with provided logger, service, and server configuration.
// It validates the inputs and returns an error if any essential component is missing.
func New(
	logger transport.Logger,
	service transport.Service,
	serverConfig *ServerConfig,
) (*HTTPServer, error) {
	if logger == nil {
		return nil, merror.ErrInvalidParameter("logger")
	}

	if service == nil {
		return nil, merror.ErrInvalidParameter("service")
	}

	if serverConfig == nil {
		return nil, merror.ErrInvalidParameter("serverConfig")
	}

	return &HTTPServer{
		Logger:       logger,
		Service:      service,
		ServerConfig: serverConfig,
	}, nil
}

// GetRouter returns the router of the http server transport layer
// with the handlers registered
func (s *HTTPServer) GetRouter() *mux.Router {
	var router = mux.NewRouter()
	router.HandleFunc("/estimate/deliverytime", s.GetDelivery).Methods(http.MethodPost)
	return router
}

// Start starts the http server on specific port with handlers registered
func (h *HTTPServer) Start() error {
	router := h.GetRouter()
	c := getCors(transport.GetConfig())
	handler := c.Handler(router)

	server := &http.Server{
		Addr:              h.ServerConfig.Port,
		ReadHeaderTimeout: time.Duration(h.ServerConfig.TimeOut) * time.Second,
		Handler:           handler,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(fmt.Errorf("listener_and_serve: %w", err))
	}
	return nil
}

func getCors(httpConfig *transport.Config) *cors.Cors {
	return cors.New(cors.Options{
		AllowedHeaders: httpConfig.Headers,
		AllowedOrigins: httpConfig.Hosts,
		AllowedMethods: httpConfig.Methods,
	})
}
