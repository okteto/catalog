package api

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

type HealthClient struct {
	URL        string
	HttpClient http.Client
}

func (h HealthClient) Get() (ServiceHealth, error) {
	resp, err := h.HttpClient.Get(h.URL)
	if err != nil {
		return ServiceHealth{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("non-200 health response: %d", resp.StatusCode)
		return ServiceHealth{}, err
	}

	var result ServiceHealth
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ServiceHealth{}, err
	}
	return result, nil
}

// type OwnerRegistrationClient struct {
// 	URL string
// }

// type ServiceRegistrationClient struct {
// 	URL string
// }

type APIHandler struct {
	HealthClient HealthClient
}

func (a APIHandler) Handle(w http.ResponseWriter, r *http.Request) {
	serviceHealth, err := a.HealthClient.Get()
	if err != nil {
		http.Error(w, fmt.Sprintf("health: %+v\n", err), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.MarshalIndent(&serviceHealth, "", "\t")
	if err != nil {
		http.Error(w, fmt.Sprintf("json: %+v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
