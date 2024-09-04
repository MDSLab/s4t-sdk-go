package boards

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MIKE9708/s4t-sdk-go/pkg"
	"github.com/MIKE9708/s4t-sdk-go/pkg/utils"
	"k8s.io/apimachinery/pkg/runtime"
)
// +kubebuilder:object:generate=true
type Board struct {
	Uuid string `json:"uuid"`
	Code string `json:"code"`
	Status string `json:"status"`
	Name string `json:"name"`
	Type string `json:"type"`
	Agent string `json:"agent"`
	Wstunip string `json:"wstun_ip,omitempty"`
	Session string `json:"session"`
	Fleet runtime.RawExtension `json:"fleet,omitempty"`
	//interface{} `json:"fleet"`
	LRversion string `json:"lr_version"`
	Connectivity Connectivity `json:"connectivity"`
	Links []Link `json:"links,omitempty"`
	Location []Location  `json:"location"`
}
func (b *Board) Keys() []string {
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

type Connectivity struct {
    Iface   string `json:"iface"`
    LocalIP string `json:"local_ip"`
    MAC     string `json:"mac"`
}
func (c Connectivity) MarshalJSON() ([]byte, error) {
	if c == (Connectivity{}) {
		return []byte("{}"), nil
	}
	type ConnectivityAlias Connectivity
	return json.Marshal(ConnectivityAlias(c))
}


type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type Location struct {
	Longitude  string      `json:"longitude"`
	Latitude   string      `json:"latitude"`
	Altitude   string      `json:"altitude"`
	UpdatedAt  []byte  `json:"updated_at"`
}

type Action struct {
	ServiceAction string `json:"service_action"`
}

type Sensors struct {
	Name string
}

func (b *Board)ListBoards(client *s4t.Client) ([]Board, error) {
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
		Boards []Board `json:"boards"`
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	

	return result.Boards, nil
}


func (b *Board)GetBoardDetail(client *s4t.Client) (*Board, error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/boards/" + b.Uuid, nil)

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

	result := Board{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %v", err)
    }

    return &result, nil
}


func (b *Board)GetBoardConf(client *s4t.Client) ([]byte, error){
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/boards/" + b.Uuid + "/conf", nil)

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

	return body, nil
}


func (b *Board)getSensors(client s4t.Client) (*Sensors, error) {
	req, err := http.NewRequest("GET", client.Endpoint + "/v1/boards/sensors/", nil)

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

	result := Sensors{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

    return &result, nil
}

func (b *Board)getBoardPosHistory(client s4t.Client) (interface {}, error){
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/boards/" + b.Uuid + "/position", nil)

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

func (b *Board)getBoardConfFile(client s4t.Client) (string,error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + client.Port + "/v1/boards/" + b.Uuid + "/conf", nil)

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


func (b *Board)DeleteBoard(client *s4t.Client) error {
	req, err := http.NewRequest("DELETE", client.Endpoint + ":" + client.Port + "/v1/boards/" + b.Uuid, nil)
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
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	return nil


}


func (b *Board)CreateBoard(client *s4t.Client) (*Board, error) {
	jsonBody, err := json.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("POST", client.Endpoint + ":" + client.Port + "/v1/boards/", bytes.NewBuffer(jsonBody))
	
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
	
	result := Board{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return b, nil
}

func (b *Board)AddNewPosition(client *s4t.Client, position Location) error {
	jsonBody, err := json.Marshal(position)
	
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("POST", client.Endpoint + ":" + client.Port + "/v1/boards/" + b.Uuid + "/position", bytes.NewBuffer(jsonBody))
	
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

func (b *Board)PatchBoard(client *s4t.Client, data map[string]interface{}) (*Board, error) {
	board := Board{}
	board_keys := board.Keys()
	compare_result := utils.CompareFields(data, board_keys)

	if !compare_result {
		return nil, fmt.Errorf("Error keys not correct")
		
	}

	jsonBody, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("PATCH", client.Endpoint + ":" + client.Port + "/v1/boards/" + b.Uuid, bytes.NewBuffer(jsonBody))
	
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
	
	result := Board{}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil
}


func (b *Board)PerformBoardAction(client *s4t.Client, action map[string] interface{}) error {
	jsonBody, err := json.Marshal(action)

	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", client.Endpoint + ":" + client.Port + "/v1/boards/" + b.Uuid + "/action", bytes.NewBuffer(jsonBody))
	
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
