package services

import (
	"fmt"
	"encoding/json"
	"bytes"
	"io"
	"net/http"
	"s4t-sdk-module/pkg"
	"s4t-sdk-module/pkg/utils"
	"s4t-sdk-module/pkg/api/boards"
)

type Service struct{
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name"`
	Project string `json:"project,omitempty"`
	Port uint `json:"port"`
	Protocol string `json:"protocol"`
	Links []boards.Link `json:"links,omitempty"`
}

func (b Service) Keys() []string {
    return  []string{
		"uuid", "code", 
		"status", "name", 
		"type", "agent", 
		"wstpun_ip","session",
		"fleet","lr_version",
		"connectivity","links",
		"location",
	}
}

func GetServices(client *s4t.Client) ([]Service, error){
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/services/", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Services []Service `json:"services"`
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	
	}

    return result.Services, nil
}

func CreateService(client *s4t.Client, service Service) (*Service,error) {
	jsonBody, err := json.Marshal(service)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("POST", client.Endpoint + ":" + client.Port + "/v1/services/", bytes.NewBuffer(jsonBody))
	
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	var result = Service{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	
	}
	return &result, nil
}

func PatchService(client *s4t.Client, service_id string, data map[string] interface{}) (*Service, error) {
	service := Service{}
	service_keys := service.Keys()
	compare_result := utils.CompareFields(data, service_keys)

	if !compare_result {
		return nil, fmt.Errorf("Error keys not correct")
		
	}

	jsonBody, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("PATCH", client.Endpoint + ":" + client.Port + "/v1/services/" + service_id, bytes.NewBuffer(jsonBody))
	
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	
	result := Service{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil
}


func DeleteService(client *s4t.Client, service_id string) error {
	req, err := http.NewRequest("DELETE", client.Endpoint + ":" + client.Port + "/v1/services/" + service_id, nil)
	
	if err != nil {
		return fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	return nil
}


func GetBoardExposedServices(client *s4t.Client, board_id string) ([]Service, error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/boards/" + board_id + "/services", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Services []Service `json:"exposed"`
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	
	}

    return result.Services, nil
}
