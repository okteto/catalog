package health

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HealthResult is the health result from a health check at the timestamp.
type HealthResult struct {
	Healthy   bool   `json:"healthy"`
	Timestamp string `json:"timestamp"`
}

// ServiceHealth provides a health result for a particular service.
type ServiceHealth struct {
	ServiceID     string         `json:"service_id"`
	HealthResults []HealthResult `json:"health_results"`
}

// HealthClient is an interface for obtaining service health data.
type HealthClient interface {
	// Get provides all known health data for services registered with the catalog.
	Get() []ServiceHealth
}

// Handler provides http handlers for providing service health data.
type Handler struct {
	HealthClient HealthClient
}

// Handle fulfills the http api for obtaining service health data.
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	results := h.HealthClient.Get()

	jsonBytes, err := json.MarshalIndent(&results, "", "\t")
	if err != nil {
		http.Error(w, fmt.Sprintf("json: %+v\n", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
