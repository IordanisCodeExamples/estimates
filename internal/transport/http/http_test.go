package transporthttp_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	merror "github.com/junkd0g/neji"
	"github.com/stretchr/testify/assert"

	transporthttpmock "github.com/junkd0g/estimates/internal/mocks/transport"
	"github.com/junkd0g/estimates/internal/transport"
	transporthttp "github.com/junkd0g/estimates/internal/transport/http"
)

type mocks struct {
	ctx     context.Context
	service *transporthttpmock.MockService
	logger  *transporthttpmock.MockLogger
}

func getMocks(t *testing.T) *mocks {
	t.Helper()
	ctrl := gomock.NewController(t)
	service := transporthttpmock.NewMockService(ctrl)
	logger := transporthttpmock.NewMockLogger(ctrl)
	return &mocks{
		service: service,
		logger:  logger,
	}
}

func Test_New(t *testing.T) {
	mocks := getMocks(t)
	tests := []struct {
		name          string
		service       transport.Service
		logger        transport.Logger
		serverConfig  *transporthttp.ServerConfig
		expectedError error
	}{
		{
			name:          "Creates successfully a http object",
			service:       mocks.service,
			logger:        mocks.logger,
			serverConfig:  &transporthttp.ServerConfig{},
			expectedError: nil,
		},
		{
			name:          "Fails to create a http object due to missing service",
			service:       nil,
			logger:        mocks.logger,
			serverConfig:  &transporthttp.ServerConfig{},
			expectedError: merror.ErrInvalidParameter("service"),
		},
		{
			name:          "Fails to create a http object due to missing logger",
			service:       mocks.service,
			logger:        nil,
			serverConfig:  &transporthttp.ServerConfig{},
			expectedError: merror.ErrInvalidParameter("logger"),
		},
		{
			name:          "Fails to create a http object due to missing serverConfig",
			service:       mocks.service,
			logger:        mocks.logger,
			serverConfig:  nil,
			expectedError: merror.ErrInvalidParameter("serverConfig"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := transporthttp.New(tc.logger, tc.service, tc.serverConfig)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
