package read_config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

type AuthRequest_1 struct {
	Identity Identity `json:"identity"`
	Scope Scope `json:"scope"`
}

type Identity struct {
	Methods []string `json:"methods"`
	Password Password `json:"password"`

}

type Password struct {
	User User `json:"user"`
}

type User struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Domain Id `json:"domain"`
}

type Id struct {
	Id string`json:"id"`
}

type Scope struct {
	Project Project `json:"project"`

}

type Project struct {
	Name string `json:"name"`
	Domain Id `json:"domain"`
}

type ConfigData struct {
	S4tAuthData struct {
		Ip string `yaml:"ip"`
		Port string `yaml:"port"`
		AuthPort string `yaml:"auth_port"`
		Token string `yaml:"temp_auth_token"` 
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"s4t-auth-data"`
	Domain struct {
		DomainName string `yaml:"name"` 
	} `yaml:"domain"`
}

func ReadConfiguration() (*ConfigData,error) {
	config := &ConfigData{}
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

func FormatAuthRequ(name string, password string, id string) *AuthRequest_1 {
	auth_req := AuthRequest_1{ // HORRIBLE
		Identity: Identity{
			Methods: []string{"password"},
			Password: Password{
				User: User{
					Name: name,
					Password: password,
					Domain: Id{
						Id: id,	
					},
				},
			},
		},
		Scope: Scope{
			Project: Project{
				Name: name,
				Domain: Id{
					Id: id,
				},
			},	
		},	
	} 
	return &auth_req
}
