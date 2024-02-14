package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/junkd0g/estimates/internal/service"
	client "github.com/junkd0g/go-client-cintworks"
	"github.com/stretchr/testify/assert"
)

func TestGetEstimates(t *testing.T) {
	t.Run("success for 1 day", func(t *testing.T) {
		mocks := getMocks(t)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{
				Feasibility: 0.5,
			}, nil).
			Times(1)

		service, err := service.New(mocks.cintAPI)
		assert.NoError(t, err)

		// Call the GetEstimates method with the surveySettings
		success, estimate, err := service.GetEstimates(context.Background(), 100, 30)
		assert.True(t, success)
		assert.Equal(t, "1 day", estimate)
		assert.NoError(t, err)
	})

	t.Run("success for 1 week", func(t *testing.T) {
		mocks := getMocks(t)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{
				Feasibility: 0.9,
			}, nil).
			Times(1)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{
				Feasibility: 0.6,
			}, nil).
			Times(1)

		service, err := service.New(mocks.cintAPI)
		assert.NoError(t, err)

		// Call the GetEstimates method with the surveySettings
		success, estimate, err := service.GetEstimates(context.Background(), 100, 30)
		assert.True(t, success)
		assert.Equal(t, "1 week", estimate)
		assert.NoError(t, err)
	})

	t.Run("success for 1 month", func(t *testing.T) {
		mocks := getMocks(t)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{
				Feasibility: 1.9,
			}, nil).
			Times(1)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{
				Feasibility: 0.9,
			}, nil).
			Times(1)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{
				Feasibility: 0.3,
			}, nil).
			Times(1)

		service, err := service.New(mocks.cintAPI)
		assert.NoError(t, err)

		// Call the GetEstimates method with the surveySettings
		success, estimate, err := service.GetEstimates(context.Background(), 100, 30)
		assert.True(t, success)
		assert.Equal(t, "1 month", estimate)
		assert.NoError(t, err)
	})

	t.Run("success on cannot be completed within 1 month", func(t *testing.T) {
		mocks := getMocks(t)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{
				Feasibility: 1.3,
			}, nil).
			Times(1)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{
				Feasibility: 0.9,
			}, nil).
			Times(1)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{
				Feasibility: 0.8,
			}, nil).
			Times(1)

		service, err := service.New(mocks.cintAPI)
		assert.NoError(t, err)

		// Call the GetEstimates method with the surveySettings
		success, estimate, err := service.GetEstimates(context.Background(), 100, 30)
		assert.False(t, success)
		assert.Equal(t, "The project cannot be completed within 1 month.", estimate)
		assert.NoError(t, err)
	})

	t.Run("error on GetFeasibilityEstimates", func(t *testing.T) {
		mocks := getMocks(t)

		mocks.cintAPI.
			EXPECT().
			GetFeasibilityEstimates(gomock.Any()).
			Return(client.SurveyResponse{}, assert.AnError).
			Times(1)

		service, err := service.New(mocks.cintAPI)
		assert.NoError(t, err)

		// Call the GetEstimates method with the surveySettings
		success, estimate, err := service.GetEstimates(context.Background(), 100, 30)
		assert.False(t, success)
		assert.Equal(t, "", estimate)
		assert.Error(t, err)
	})
}
