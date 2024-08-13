package boards

import (
	"s4t-sdk-module/pkg"
	"net/http"
	"fmt"
	"encoding/json"
)

type Board struct {
	Uuid string
	Code string
	Atatus string
	Name string
	Type string
	Agent string
	Wstunip string
	Owner string
	Session string
	Project string
	Fleet string
	Mobile bool
	Lr_version string
	Connectivity string `json:"connectivity"`
	Extra string `json:"extra"`
	Config string `json:"config"`
	Links []Link
	Location Location
}

type Link struct {
	href string
	rel string
}

type Location struct {
	longitude string
	latitude string
	altitude string
}


type Sensors struct {
	name string
}

func ListBoards (client s4t.Client) ([]Board, error) {
	req, err := http.NewRequest("GET", client.Endpoint + "/v1/boards", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Boards []Board `json:"servers"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %v", err)
    }

    return result.Boards, nil
}


func getBoardDetail(client s4t.Client, board_name string) (*Board, error) {
	req, err := http.NewRequest("GET", client.Endpoint + "/v1/boards/" + board_name, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	result := Board{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %v", err)
    }

    return &result, nil
}

func getSensors(client s4t.Client) (*Sensors, error) {
	req, err := http.NewRequest("GET", client.Endpoint + "/v1/sensors/", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	result := Sensors{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %v", err)
    }

    return &result, nil
}


func getBoardPosHistory(client s4t.Client, board_name string) (interface {}, error){
	req, err := http.NewRequest("GET", client.Endpoint + "/v1/boards/" + board_name + "/position", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %v", err)
    }

    return result, nil
}


func getBoardConfFile(client s4t.Client, board_name string) (string,error) {
	req, err := http.NewRequest("GET", client.Endpoint + "/v1/boards/" + board_name + "/conf", nil)

	if err != nil {
		return "", fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", client.AuthToken)

	resp, err := client.HTTPClient.Do(req)

	if err != nil {
		return "", fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}
	// TO BE CHANGED
	return "", nil
}
