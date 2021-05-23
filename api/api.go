package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/okteto/catalog/health"
)

type HealthClient struct {
	URL        string
	HttpClient http.Client
}

func (h HealthClient) Get() (health.ServiceHealth, error) {
	resp, err := h.HttpClient.Get(h.URL)
	if err != nil {
		return health.ServiceHealth{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("non-200 health response: %d", resp.StatusCode)
		return health.ServiceHealth{}, err
	}

	var result health.ServiceHealth
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return health.ServiceHealth{}, err
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
