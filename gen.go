// Package gen provides code generation utilities for the project.
package gen

//go:generate mockgen -package servicemock -destination internal/mocks/service/service.go -source=internal/service/service.go CintAPI
//go:generate mockgen -package transportmock -destination internal/mocks/transport/transport.go -source=internal/transport/transport.go Service Logger
