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

type Owner struct {
	ID         string
	Name       string
	ServiceIDs []string
}

type Service struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	HealthAddr string `json:"health_addr"`
}

type HealthClient struct {
	URL        string
	HTTPClient http.Client
}

func (h HealthClient) Get() ([]ServiceHealth, error) {
	resp, err := h.HTTPClient.Get(h.URL)
	if err != nil {
		return []ServiceHealth{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("non-200 health response: %d", resp.StatusCode)
		return []ServiceHealth{}, err
	}

	var result []ServiceHealth
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []ServiceHealth{}, err
	}
	return result, nil
}

type OwnerRegistrationClient struct {
	URL        string
	HTTPClient http.Client
}

func (o OwnerRegistrationClient) Get() ([]Owner, error) {
	resp, err := o.HTTPClient.Get(o.URL)
	if err != nil {
		return []Owner{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("non-200 owner response: %d", resp.StatusCode)
		return []Owner{}, err
	}

	var result []Owner
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []Owner{}, err
	}
	return result, nil
}

type ServiceRegistrationClient struct {
	URL        string
	HTTPClient http.Client
}

func (s ServiceRegistrationClient) Get() ([]Service, error) {
	resp, err := s.HTTPClient.Get(s.URL)
	if err != nil {
		return []Service{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("non-200 service response: %d", resp.StatusCode)
		return []Service{}, err
	}

	var result []Service
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []Service{}, err
	}
	return result, nil
}

type CatalogDataEntry struct {
	ServiceName string         `json:"service_name"`
	OwnerID     string         `json:"owner_id"`
	OwnerName   string         `json:"owner_name"`
	HealthData  []HealthResult `json:"health_data`
}

type CatalogData map[string]CatalogDataEntry

type APIHandler struct {
	HealthClient              HealthClient
	OwnerRegistrationClient   OwnerRegistrationClient
	ServiceRegistrationClient ServiceRegistrationClient
}

func (a APIHandler) Handle(w http.ResponseWriter, r *http.Request) {
	serviceHealth, err := a.HealthClient.Get()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("error retrieving health data: %+v\n", err),
			http.StatusInternalServerError,
		)
		return
	}

	registeredOwners, err := a.OwnerRegistrationClient.Get()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("error retrieving owner data: %+v\n", err),
			http.StatusInternalServerError,
		)
		return
	}

	registerdServices, err := a.ServiceRegistrationClient.Get()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("error retrieving service data: %+v\n", err),
			http.StatusInternalServerError,
		)
		return
	}

	catalogData := make(CatalogData)
	catalogData = populateServices(catalogData, registerdServices)
	catalogData = populateOwners(catalogData, registeredOwners)
	catalogData = populateHealth(catalogData, serviceHealth)

	jsonBytes, err := json.MarshalIndent(&catalogData, "", "\t")
	if err != nil {
		http.Error(w, fmt.Sprintf("json: %+v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func populateServices(data CatalogData, services []Service) CatalogData {
	for _, service := range services {
		data[service.ID] = CatalogDataEntry{
			ServiceName: service.Name,
		}
	}
	return data
}

func populateOwners(data CatalogData, owners []Owner) CatalogData {
	for _, owner := range owners {
		for _, serviceID := range owner.ServiceIDs {
			entry := data[serviceID]
			entry.OwnerID = owner.ID
			entry.OwnerName = owner.Name
			data[serviceID] = entry
		}
	}
	return data
}

func populateHealth(data CatalogData, health []ServiceHealth) CatalogData {
	for _, serviceHealth := range health {
		entry := data[serviceHealth.ServiceID]
		entry.HealthData = serviceHealth.HealthResults
		data[serviceHealth.ServiceID] = entry
	}
	return data
}
