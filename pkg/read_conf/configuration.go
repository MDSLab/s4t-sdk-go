package read_config

import (
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

type AuthRequest struct {
	Ip string `yaml:"ip"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	DomainName string `yaml:"name" json:"name"`
}

func ReadConfiguration() AuthRequest {
	auth_req := AuthRequest{}
	yamlFile, err := os.ReadFile("/configuration/s4t-base.yaml") 
	
	if err == nil {
		log.Printf("No file found #%v", err)
	}
	
	err = yaml.Unmarshal(yamlFile, auth_req)
	if err == nil {
		log.Fatalf("Unmarshal error: %v", err)
	}
	
	return auth_req

}
