package plugins

import (
	"github.com/MIKE9708/s4t-sdk-go/pkg/api/data/generic"
	"k8s.io/apimachinery/pkg/runtime"
)

type PluginReq struct {
	Name       string               `json:"name"`
	Parameters runtime.RawExtension `json:"parameters,omitempty"`
	Code       string               `json:"code"`
	Version    string               `json:"version,omitempty"`
}

func (b *PluginReq) Keys() []string {
	return []string{
		"name", "parameters",
		"code", "version",
	}
}

type Plugin struct {
	UUID       string               `json:"uuid,omitempty"`
	Name       string               `json:"name"`
	Public     bool                 `json:"public"`
	Code       string               `json:"code"`
	Parameters runtime.RawExtension `json:"parameters,omitempty"`
	Version    string               `json:"version,omitempty"`
	Owner      string               `json:"owner"`
	Callable   bool                 `json:"callable"`
	Links      []generic.Link       `json:"links,omitempty"`
}
