package s4t

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	boards "github.com/MIKE9708/s4t-sdk-go/pkg/api/data/board"
	services "github.com/MIKE9708/s4t-sdk-go/pkg/api/data/service"
	"github.com/MIKE9708/s4t-sdk-go/pkg/utils"
)

func (client *Client) GetServices() ([]services.Service, error) {
	req, err := http.NewRequest("GET", client.Endpoint+":"+client.Port+"/v1/services/", nil)

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
		Services []services.Service `json:"services"`
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)

	}

	return result.Services, nil
}

func (client *Client) CreateService(service services.Service) (*services.Service, error) {
	jsonBody, err := json.Marshal(service)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)

	}
	req, err := http.NewRequest("POST", client.Endpoint+":"+client.Port+"/v1/services/", bytes.NewBuffer(jsonBody))

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
	var result = services.Service{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)

	}
	return &result, nil
}

func (client *Client) PatchService(service_id string, data map[string]interface{}) (*services.Service, error) {
	service := services.Service{}
	service_keys := service.Keys()
	compare_result := utils.CompareFields(data, service_keys)

	if !compare_result {
		return nil, fmt.Errorf("Error keys not correct")

	}

	jsonBody, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)

	}
	req, err := http.NewRequest("PATCH", client.Endpoint+":"+client.Port+"/v1/services/"+service_id, bytes.NewBuffer(jsonBody))

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

	result := services.Service{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil
}

func (client *Client) DeleteService(service_id string) error {
	req, err := http.NewRequest("DELETE", client.Endpoint+":"+client.Port+"/v1/services/"+service_id, nil)

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

func (client *Client) GetBoardExposedServices(board_id string) ([]services.Service, error) {
	req, err := http.NewRequest("GET", client.Endpoint+":"+client.Port+"/v1/boards/"+board_id+"/services", nil)

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
		Services []services.Service `json:"exposed"`
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)

	}

	return result.Services, nil
}

func (client *Client) RestoreService(board_id string) error {
	req, err := http.NewRequest("GET", client.Endpoint+":"+client.Port+"/v1/boards/"+board_id+"/services/restore", nil)

	if err != nil {
		return fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (client *Client) PerfomActionOnService(
	board_id string,
	service_id string, action boards.Action) error {

	jsonBody, err := json.Marshal(action)
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)

	}
	req, err := http.NewRequest("POST",
		client.Endpoint+":"+client.Port+"/v1/boards/"+board_id+"/services"+service_id+"/action",
		bytes.NewBuffer(jsonBody))

	if err != nil {
		return fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()

	return nil
}
