package main

import (
	"fmt"
	"log"
	"s4t-sdk-module/internal/auth"
	"s4t-sdk-module/pkg"
	"s4t-sdk-module/pkg/read_conf"
)

func main () {
	auth_req := read_config.ReadConfiguration()

	client := s4t.NewClient("http://" + auth_req.Ip)
	
	resp, err := auth.Authenticate(client, &auth_req)
	
	fmt.Printf("%v", resp)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
	}

	// computeClient := compute.NewComputeClient(client.HTTPClient, client.AuthToken, client.Endpoint + "/compute")

	// _, err := computeClient.ListServers()

	// if err == nil {
		// log.Fatal("Failed to list servers: %v", err)
	//}

}
