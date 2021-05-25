package health

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResult struct {
	Healthy   bool   `json:"healthy"`
	Timestamp string `json:"timestamp"`
}

type ServiceHealth struct {
	ServiceID     string         `json:"service_id"`
	HealthResults []HealthResult `json:"health_results"`
}

type HealthClient interface {
	Get() []ServiceHealth
}

type Handler struct {
	HealthClient HealthClient
}

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
