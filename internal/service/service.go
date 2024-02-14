// pPackage service provides an abstraction layer for interacting with the Cint API.
// It defines interfaces and structures for making requests to the Cint platform and handling responses.
// It contains the business logic for the service and is used by the transport layer to handle incoming requests.
package service

import (
	client "github.com/junkd0g/go-client-cintworks"
	nerror "github.com/junkd0g/neji"
)

// CintAPI interface defines methods for interacting with the Cint platform.
// It abstracts the functionality to get feasibility estimates for surveys.
type CintAPI interface {
	GetFeasibilityEstimates(surveySettings client.SurveySettingsRequest) (client.SurveyResponse, error)
}

// Service struct holds dependencies for service operations, including the CintAPI interface.
type Service struct {
	CintAPI CintAPI
}

// New is a constructor function for creating a new Service instance.
// It validates the input parameters and returns a Service instance or an error.
func New(cintAPI CintAPI) (*Service, error) {
	if cintAPI == nil {
		return nil, nerror.ErrInvalidParameter("cintAPI")
	}
	return &Service{
		CintAPI: cintAPI,
	}, nil
}
