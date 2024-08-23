package read_config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AuthRequest struct {
	S4tAuthData struct {
		Ip string `yaml:"ip"`
		Port string `yaml:"port"`
		Token string `yaml:"temp_auth_token"` 
		Username string `yaml:"username" json:"username"`
		Password string `yaml:"password" json:"password"`
	} `yaml:"s4t-auth-data"`
	Domain struct {
		DomainName string `yaml:"name" json:"name"`
	} `yaml:"domain"`
}

func ReadConfiguration() (*AuthRequest,error) {
	config := &AuthRequest{}
	file, err := os.Open("../../configuration/s4t-base.yaml") 
	
	if err != nil {
		return nil, fmt.Errorf("Error opening the file #%v", err)
	}
	defer file.Close()
	
	d := yaml.NewDecoder(file)
	
	if err := d.Decode(&config); err != nil {		
		return nil, fmt.Errorf("Error decoding YAML: %v", err)
	}
	
    // fmt.Printf("Ip: %s\n", config.S4tAuthData.Ip)
    // fmt.Printf("User: %s\n", config.S4tAuthData.Username)
    // fmt.Printf("Domain: %s\n", config.Domain.DomainName)
	
	return config, nil


}
