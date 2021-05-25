package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	HealthAddr string `json:"health_addr"`
}

var registeredServices = []Service{
	{
		ID:         "3d8201b4-d152-4d84-b8bc-eb87817a617e",
		Name:       "Client API",
		HealthAddr: "client-api.okteto:8081/healthz",
	},
	{
		ID:         "a41a7cb2-6ab1-489d-9dab-ff008e4ce34e",
		Name:       "Webhook Registry",
		HealthAddr: "webhook-registry.okteto:8081/healthz",
	},
	{
		ID:         "028d4690-f9b8-4cb2-9edb-a44d831dbed3",
		Name:       "User Data",
		HealthAddr: "user-data.okteto:8081/healthz",
	},
	{
		ID:         "d2647c9b-2e15-4978-a97f-2bcad9126f57",
		Name:       "Event History",
		HealthAddr: "event-history.okteto:8081/healthz",
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.MarshalIndent(&registeredServices, "", "\t")
	if err != nil {
		http.Error(w, fmt.Sprintf("json: %+v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
