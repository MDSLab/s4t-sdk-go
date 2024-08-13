package compute

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type ComputeClient struct {
	client *http.Client
	token string
	url string
}


func NewComputeClient (client *http.Client, token, url string) *ComputeClient {
	return &ComputeClient {
		client: client,
		token: token,
		url: url,
	}

}

type Server struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`

}

func (c *ComputeClient) ListServers () ([]Server, error) {
	req, err := http.NewRequest("GET", c.url + "/servers", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	req.Header.Set("X-Auth-Token", c.token)

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Servers []Server `json:"servers"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %v", err)
    }

    return result.Servers, nil
} 
