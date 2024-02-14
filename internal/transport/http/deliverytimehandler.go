package transporthttp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	nerror "github.com/junkd0g/neji"

	internallogger "github.com/junkd0g/estimates/internal/logger"
)

func (s *HTTPServer) GetDelivery(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")

	jsonBody, status := s.getDelivery(ctx, request)
	writer.WriteHeader(status)
	_, err := writer.Write(jsonBody)
	if err != nil {
		s.Logger.Error(ctx, err.Error())
	}
}

func (h *HTTPServer) getDelivery(ctx context.Context, request *http.Request) ([]byte, int) {
	h.Logger.Info(ctx, "getDelivery started")

	var req Request

	b, err := io.ReadAll(request.Body)
	if err != nil {
		h.Logger.Error(ctx, "error_reading_body_transporthttp_getdelivery", internallogger.LogField{"error": err.Error()})
		resp, _ := nerror.SimpeErrorResponseWithStatus(http.StatusBadRequest, err)
		return resp, http.StatusBadRequest
	}

	err = json.Unmarshal(b, &req)
	if err != nil {
		h.Logger.Error(ctx, "error_unmarshalling_body_transporthttp_getdelivery", internallogger.LogField{"error": err.Error()})
		resp, _ := nerror.SimpeErrorResponseWithStatus(http.StatusBadRequest, err)
		return resp, http.StatusBadRequest
	}

	isFeasible, message, err := h.Service.GetEstimates(ctx, req.Limit, req.LengthOfInterview)
	if err != nil {
		h.Logger.Error(ctx, "error_fixturesbydate_transporthttp_getdelivery", internallogger.LogField{"error": err.Error()})
		resp, _ := nerror.SimpeErrorResponseWithStatus(http.StatusBadRequest, err)
		return resp, http.StatusBadRequest
	}

	res := Response{
		Feasible:     isFeasible,
		DeliveryTime: message,
	}

	jsonBody, err := json.Marshal(res)
	if err != nil {
		h.Logger.Error(ctx, "error_marshaling_response_transporthttp_getdelivery", internallogger.LogField{"error": err.Error()})
		resp, _ := nerror.SimpeErrorResponseWithStatus(http.StatusInternalServerError, err)
		return resp, http.StatusInternalServerError
	}

	h.Logger.Info(ctx, "getDelivery ended")

	return jsonBody, http.StatusOK
}
