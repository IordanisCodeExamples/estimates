package transporthttp_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	transporthttp "github.com/junkd0g/estimates/internal/transport/http"
)

func TestGetDeliveryHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mocks := getMocks(t)
		// Create an instance of the Request struct with sample data
		requestData := transporthttp.Request{
			Limit:             100, // Example limit
			LengthOfInterview: 30,  // Example length of interview in minutes
		}

		// Marshal the requestData into JSON
		requestBodyBytes, err := json.Marshal(requestData)
		if err != nil {
			t.Fatalf("Failed to marshal request body: %v", err)
		}
		requestBody := bytes.NewReader(requestBodyBytes)

		mocks.service.
			EXPECT().
			GetEstimates(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(true, "1 month", nil)

		mocks.logger.
			EXPECT().
			Info(gomock.Any(), "getDelivery started")

		mocks.logger.
			EXPECT().
			Info(gomock.Any(), "getDelivery ended")

		server, err := transporthttp.New(
			mocks.logger,
			mocks.service,
			&transporthttp.ServerConfig{
				Port:    ":8080",
				TimeOut: 10,
			},
		)
		assert.NoError(t, err)

		// Adjust the endpoint and method for the POST request, including the JSON body
		req := httptest.NewRequest("POST", "/estimate/deliverytime", requestBody)
		// Set the Content-Type header to application/json
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		server.GetDelivery(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("error_reading_body", func(t *testing.T) {
		mocks := getMocks(t)
		requestBody := bytes.NewReader([]byte("{invalid json"))

		mocks.logger.EXPECT().
			Info(gomock.Any(), "getDelivery started")

		mocks.logger.EXPECT().
			Error(gomock.Any(), gomock.Eq("error_unmarshalling_body_transporthttp_getdelivery"), gomock.Any()).Times(1)

		server, err := transporthttp.New(
			mocks.logger,
			mocks.service,
			&transporthttp.ServerConfig{
				Port:    ":8080",
				TimeOut: 10,
			},
		)
		assert.NoError(t, err)

		req := httptest.NewRequest("POST", "/estimate/deliverytime", requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		server.GetDelivery(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code) // Expecting bad request status
	})

	t.Run("error_unmarshalling_body", func(t *testing.T) {
		mocks := getMocks(t)
		requestBody := bytes.NewReader([]byte("{\"limit\": \"invalid\", \"lengthOfInterview\": 30}")) // Invalid JSON for the expected struct

		mocks.logger.EXPECT().Info(gomock.Any(), "getDelivery started")
		mocks.logger.EXPECT().
			Error(gomock.Any(), gomock.Eq("error_unmarshalling_body_transporthttp_getdelivery"), gomock.Any()).Times(1)

		server, err := transporthttp.New(
			mocks.logger,
			mocks.service,
			&transporthttp.ServerConfig{
				Port:    ":8080",
				TimeOut: 10,
			},
		)
		assert.NoError(t, err)

		req := httptest.NewRequest("POST", "/estimate/deliverytime", requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		server.GetDelivery(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("service_error", func(t *testing.T) {
		mocks := getMocks(t)
		requestData := transporthttp.Request{
			Limit:             100,
			LengthOfInterview: 30,
		}
		requestBodyBytes, err := json.Marshal(requestData)
		assert.NoError(t, err)
		requestBody := bytes.NewReader(requestBodyBytes)

		mocks.service.EXPECT().
			GetEstimates(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(false, "", errors.New("service error")) // Simulate service error

		mocks.logger.EXPECT().Info(gomock.Any(), "getDelivery started")
		mocks.logger.EXPECT().
			Error(gomock.Any(), gomock.Eq("error_fixturesbydate_transporthttp_getdelivery"), gomock.Any()).Times(1)

		server, err := transporthttp.New(
			mocks.logger,
			mocks.service,
			&transporthttp.ServerConfig{
				Port:    ":8080",
				TimeOut: 10,
			},
		)
		assert.NoError(t, err)

		req := httptest.NewRequest("POST", "/estimate/deliverytime", requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		server.GetDelivery(w, req)

		assert.NotEqual(t, http.StatusOK, w.Code) // Expecting an error status code, not OK
	})

}
