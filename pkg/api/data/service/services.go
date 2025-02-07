package services

import (
	"github.com/MIKE9708/s4t-sdk-go/pkg/api/data/generic"
)

type Service struct {
	Uuid     string         `json:"uuid,omitempty"`
	Name     string         `json:"name"`
	Project  string         `json:"project,omitempty"`
	Port     uint           `json:"port"`
	Protocol string         `json:"protocol"`
	Links    []generic.Link `json:"links,omitempty"`
}

func (b *Service) Keys() []string {
	return []string{
		"uuid", "code",
		"status", "name",
		"type", "agent",
		"wstpun_ip", "session",
		"fleet", "lr_version",
		"connectivity", "links",
		"location",
	}
}
