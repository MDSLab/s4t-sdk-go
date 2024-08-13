package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"s4t-sdk-module/pkg"
	read_config "s4t-sdk-module/pkg/read_conf"
)

type AuthResponse struct {
	Token struct {
		IssuedAt string `json:"issued_at"`
		ExpiresAt string `json:"expires_at"`
		ID string `json:"id"`
	} `json:"token"`

}


func Authenticate (client *s4t.Client, auth_req *read_config.AuthRequest) (string, error) {

	jsonData, err := json.Marshal(&auth_req)

	if err != nil {
		return "",fmt.Errorf("Error marshaling to JSON: %d\n", err)
		
	}	
	
	resp, err := http.Post(client.Endpoint + "/v1/auth/", "application/json", bytes.NewBuffer(jsonData))
	
	if err != nil {
		return "", fmt.Errorf("authentication request failed: %v", err)
	}
	
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusCreated {
		return "",fmt.Errorf("Authentication failed with status code: %d", resp.StatusCode)
	}

	token := resp.Header.Get("X-Subject-Token")
	
	if token == "" {
		return "", fmt.Errorf("No token found in the response")
	}
	
	fmt.Printf("%v",resp)

	return token, nil
}
