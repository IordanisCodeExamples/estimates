package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	nerror "github.com/junkd0g/neji"
	"github.com/stretchr/testify/assert"

	servicemock "github.com/junkd0g/estimates/internal/mocks/service"
	"github.com/junkd0g/estimates/internal/service"
)

type mocks struct {
	cintAPI *servicemock.MockCintAPI
}

func getMocks(t *testing.T) *mocks {
	t.Helper()
	ctrl := gomock.NewController(t)
	cintAPI := servicemock.NewMockCintAPI(ctrl)
	return &mocks{
		cintAPI: cintAPI,
	}
}

func Test_New(t *testing.T) {
	mocks := getMocks(t)
	tests := []struct {
		name          string
		cintAPI       service.CintAPI
		expectedError error
	}{
		{
			name:          "Creates successfully a service object",
			cintAPI:       mocks.cintAPI,
			expectedError: nil,
		},
		{
			name:          "Fails to create a service object due to missing cintAPI",
			cintAPI:       nil,
			expectedError: nerror.ErrInvalidParameter("cintAPI"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.New(tt.cintAPI)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
