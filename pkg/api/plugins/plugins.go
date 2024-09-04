package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MIKE9708/s4t-sdk-go/pkg"
	"github.com/MIKE9708/s4t-sdk-go/pkg/api/boards"
	"github.com/MIKE9708/s4t-sdk-go/pkg/utils"
	"k8s.io/apimachinery/pkg/runtime"
)

type PluginReq struct {
	Name string  `json:"name"`
	Parameters  runtime.RawExtension `json:"parameters,omitempty"`
	Code string `json:"code"`
	Version string `json:"version,omitempty"`
}

func (b *PluginReq) Keys() []string {
    return  []string{
		"name", "parameters", 
		"code", "version", 
	}
}
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

func GetPlugin(client *s4t.Client, uuid string) (*Plugin ,error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/plugins/" + uuid, nil)

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


func CreatePlugin(client *s4t.Client, plugin PluginReq) (*Plugin, error) {
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

	return &result, nil
}

func DeletePlugin(client *s4t.Client, uuid string) error {
	req, err := http.NewRequest("DELETE", client.Endpoint + ":" + client.Port + "/v1/plugins/" + uuid, nil)
	
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

func PacthPlugin(client *s4t.Client, uuid string, data map[string] interface{}) (*Plugin, error) {
	plugin := PluginReq{}
	service_keys := plugin.Keys()
	compare_result := utils.CompareFields(data, service_keys)

	if !compare_result {
		return nil, fmt.Errorf("Error keys not correct")
		
	}

	jsonBody, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("PATCH", client.Endpoint + ":" + client.Port + "/v1/plugins/" + uuid, bytes.NewBuffer(jsonBody))
	
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
	
	result := Plugin{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil
} 

func GetBoardPlugins(client *s4t.Client, board_id string) ([]Plugin, error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/boards/"  + board_id + "/plugins", nil)

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

func InjectPLuginBoard(client *s4t.Client, board_id string, data map[string] interface{}) error {
	jsonBody, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("PUT", client.Endpoint + ":" + client.Port + "/v1/boards/" + board_id + "/plugins/", bytes.NewBuffer(jsonBody))
	
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
func GetPluginStatus(client *s4t.Client) {
}

// 405
func GetPluginsLog(client *s4t.Client) {}

func RemoveInjectedPlugin(client *s4t.Client,uuid string, board_id string) error {
	req, err := http.NewRequest("DELETE", client.Endpoint + ":" + client.Port + "/v1/boards/" + board_id + "/plugins/"  + uuid, nil)
	
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
