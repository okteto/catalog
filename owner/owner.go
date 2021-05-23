package owner

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Owner struct {
	ID         string
	Name       string
	ServiceIDs []string
}

var owners = []Owner{
	{
		ID:   "64565f85-a593-440d-a345-e782ddac4b11",
		Name: "Alice",
		ServiceIDs: []string{
			"3d8201b4-d152-4d84-b8bc-eb87817a617e",
			"a41a7cb2-6ab1-489d-9dab-ff008e4ce34e",
		},
	},
	{
		ID:   "830906d5-154c-462e-bba5-c2a1941c53d4",
		Name: "Charlotte",
		ServiceIDs: []string{
			"028d4690-f9b8-4cb2-9edb-a44d831dbed3",
		},
	},
	{
		ID:   "3571ae57-34b7-4c01-a850-a02620f0e07b",
		Name: "Colette",
		ServiceIDs: []string{
			"d2647c9b-2e15-4978-a97f-2bcad9126f57",
		},
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.MarshalIndent(&owners, "", "\t")
	if err != nil {
		http.Error(w, fmt.Sprintf("json: %+v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
