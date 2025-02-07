package s4t

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MIKE9708/s4t-sdk-go/pkg/api/data/plugin"
	"github.com/MIKE9708/s4t-sdk-go/pkg/utils"
	"io"
	"net/http"
)

func (client *Client) GetPlugins() ([]plugins.Plugin, error) {
	req, err := http.NewRequest("GET", client.Endpoint+":"+client.Port+"/v1/boards/", nil)

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
		Plugins []plugins.Plugin `json:"plugins"`
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return result.Plugins, nil
}

func (client *Client) GetPlugin(uuid string) (*plugins.Plugin, error) {
	req, err := http.NewRequest("GET", client.Endpoint+":"+client.Port+"/v1/plugins/"+uuid, nil)

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

	result := plugins.Plugin{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil

}

func (client *Client) CreatePlugin(plugin plugins.PluginReq) (*plugins.Plugin, error) {
	jsonBody, err := json.Marshal(plugin)
	fmt.Printf("%v", string(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)

	}
	req, err := http.NewRequest("POST", client.Endpoint+":"+client.Port+"/v1/plugins/", bytes.NewBuffer(jsonBody))

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

	result := plugins.Plugin{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil
}

func (client *Client) DeletePlugin(uuid string) error {
	req, err := http.NewRequest("DELETE", client.Endpoint+":"+client.Port+"/v1/plugins/"+uuid, nil)

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

func (client *Client) PacthPlugin(uuid string, data map[string]interface{}) (*plugins.Plugin, error) {
	plugin := plugins.PluginReq{}
	service_keys := plugin.Keys()
	compare_result := utils.CompareFields(data, service_keys)

	if !compare_result {
		return nil, fmt.Errorf("Error keys not correct")

	}

	jsonBody, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)

	}
	req, err := http.NewRequest("PATCH", client.Endpoint+":"+client.Port+"/v1/plugins/"+uuid, bytes.NewBuffer(jsonBody))

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

	result := plugins.Plugin{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil
}

func (client *Client) GetBoardPlugins(board_id string) ([]plugins.Plugin, error) {
	req, err := http.NewRequest("GET", client.Endpoint+":"+client.Port+"/v1/boards/"+board_id+"/plugins", nil)

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
		Plugins []plugins.Plugin `json:"plugins"`
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return result.Plugins, nil
}

func (client *Client) InjectPLuginBoard(board_id string, data map[string]interface{}) error {
	jsonBody, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)

	}
	req, err := http.NewRequest("PUT", client.Endpoint+":"+client.Port+"/v1/boards/"+board_id+"/plugins/", bytes.NewBuffer(jsonBody))

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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// 405
func (client *Client) GetPluginStatus() {
}

// 405
func (client *Client) GetPluginsLog() {}

func (client *Client) RemoveInjectedPlugin(uuid string, board_id string) error {
	req, err := http.NewRequest("DELETE", client.Endpoint+":"+client.Port+"/v1/boards/"+board_id+"/plugins/"+uuid, nil)

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
