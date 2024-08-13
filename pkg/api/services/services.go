package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"s4t-sdk-module/pkg"
)


func get_services(client s4t.Client) (json, error){
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

}

