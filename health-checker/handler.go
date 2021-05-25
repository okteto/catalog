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

// var results = []ServiceHealth{
// 	{
// 		ServiceID: "3d8201b4-d152-4d84-b8bc-eb87817a617e",
// 		HealthResults: []HealthResult{
// 			{
// 				Healthy:   true,
// 				Timestamp: "1621778498",
// 			},
// 		},
// 	},
// 	{
// 		ServiceID: "a41a7cb2-6ab1-489d-9dab-ff008e4ce34e",
// 		HealthResults: []HealthResult{
// 			{
// 				Healthy:   false,
// 				Timestamp: "1621778553",
// 			},
// 		},
// 	},
// 	{
// 		ServiceID: "028d4690-f9b8-4cb2-9edb-a44d831dbed3",
// 		HealthResults: []HealthResult{
// 			{
// 				Healthy:   true,
// 				Timestamp: "1621778592",
// 			},
// 		},
// 	},
// 	{
// 		ServiceID: "d2647c9b-2e15-4978-a97f-2bcad9126f57",
// 		HealthResults: []HealthResult{
// 			{
// 				Healthy:   true,
// 				Timestamp: "1621778628",
// 			},
// 		},
// 	},
// }

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
