package boards

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"s4t-sdk-module/pkg"
)

const BOARD_PORT = "8812"

type Board struct {
	Uuid string `json:"uuid"`
	Code string `json:"code"`
	Status string `json:"status"`
	Name string `json:"name"`
	Type string `json:"type"`
	Agent string `json:"agent"`
	Wstunip string `json:"wstun_ip,omitempty"`
	Session string `json:"session"`
	Fleet interface{} `json:"fleet"`
	LRversion string `json:"lr_version"`
	Connectivity Connectivity `json:"connectivity"`
	Links []Link `json:"links,omitempty"`
	Location []Location  `json:"location"`
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
	UpdatedAt  interface{} `json:"updated_at,omitempty"`
}


type Sensors struct {
	name string
}


func ListBoards (client *s4t.Client) ([]Board, error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + BOARD_PORT + "/v1/boards/" , nil)

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


func GetBoardDetail(client *s4t.Client, board_id string) (*Board, error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + BOARD_PORT + "/v1/boards/" + board_id, nil)

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

// NOT WORKING NO ENTRY FOUND 
func getBoardPosHistory(client s4t.Client, board_id string) (interface {}, error){
	req, err := http.NewRequest("GET", client.Endpoint + "/v1/boards/" + board_id + "/position", nil)

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


// NOT WORKING NO ENTRY FOUND 
func getBoardConfFile(client s4t.Client, board_id string) (string,error) {
	req, err := http.NewRequest("GET", client.Endpoint + ":" + BOARD_PORT + "/v1/boards/" + board_id + "/conf", nil)

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


func DeleteBoard(client *s4t.Client, board_id string) error {
	req, err := http.NewRequest("DELETE", client.Endpoint + ":" + BOARD_PORT + "/v1/boards/" + board_id, nil)
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


func CreateBoard(client *s4t.Client, board Board) (*Board, error) {
	jsonBody, err := json.Marshal(board)
	fmt.Println(string(jsonBody))	
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("POST", client.Endpoint + ":" + BOARD_PORT + "/v1/boards/", bytes.NewBuffer(jsonBody))
	
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

	return &board, nil
}


func AddNewPosition(client *s4t.Client, board_id string, position Location) error {
	jsonBody, err := json.Marshal(position)
	fmt.Println(string(jsonBody))	
	
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
		
	}
	req, err := http.NewRequest("POST", client.Endpoint + ":" + BOARD_PORT + "/v1/boards/" + board_id + "/position", bytes.NewBuffer(jsonBody))
	
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

