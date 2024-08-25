package plugins

import (
	"s4t-sdk-module/pkg"
	"s4t-sdk-module/pkg/api/boards"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Plugin struct {
    UUID     string `json:"uuid,omitempty"`
    Name     string `json:"name"`
    Public   bool   `json:"public"`
    Owner    string `json:"owner"`
    Callable bool   `json:"callable"`
    Links    []boards.Link `json:"links,omitempty"`
}


func GetPlugins(client *s4t.Client) ([]Plugin, error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/boards/" , nil)

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
	
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Plugins []Plugin `json:"plugins"`
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return result.Plugins, nil
}

func GetPlugin(client *s4t.Client, plugin_id string) (*Plugin ,error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/plugins/" + plugin_id , nil)

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
	
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	result := Plugin{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil


}


func CreatePlugin(client *s4t.Client, plugin Plugin) (*Plugin, error) {
	jsonBody, err := json.Marshal(plugin)
	fmt.Printf("%v", string(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("POST", client.Endpoint + ":" + client.Port + "/v1/plugins/", bytes.NewBuffer(jsonBody))
	
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
	
	result := Plugin{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &plugin, nil

}

