package handlers

import (
	"api_metering_system/metering"
	"encoding/json"
	"net/http"
)

// ApiHandler holds the API metering logic.
type ApiHandler struct {
	apiMetrics *metering.ApiMetrics
}

// NewApiHandler creates a new ApiHandler.
func NewApiHandler(apiMetrics *metering.ApiMetrics) *ApiHandler {
	return &ApiHandler{apiMetrics: apiMetrics}
}

// TrackEndpoint1 tracks requests to endpoint1 and responds with "OK".
func (h *ApiHandler) TrackEndpoint1(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	h.apiMetrics.Increment("/api/endpoint1")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// GetMetrics returns the current API request metrics as JSON.
func (h *ApiHandler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	metrics := h.apiMetrics.GetMetrics()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}
