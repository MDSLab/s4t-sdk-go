package boards

import (
	"encoding/json"
	"github.com/MDSLab/s4t-sdk-go/pkg/api/data/generic"
	"k8s.io/apimachinery/pkg/runtime"
	"time"
)

type Board struct {
	Uuid    string                `json:"uuid,omitempty"`
	Code    string                `json:"code"`
	Status  string                `json:"status,omitempty"`
	Name    string                `json:"name"`
	Type    string                `json:"type,omitempty"`
	Agent   string                `json:"agent,omitempty"`
	Wstunip string                `json:"wstun_ip,omitempty"`
	Session string                `json:"session,omitempty"`
	Fleet   *runtime.RawExtension `json:"fleet,omitempty"`
	//interface{} `json:"fleet"`
	LRversion    string          `json:"lr_version,omitempty"`
	Connectivity *Connectivity   `json:"connectivity,omitempty"`
	Links        []*generic.Link `json:"links,omitempty"`
	Location     []*Location     `json:"location"`
}

func (b *Board) Keys() []string {
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

type Connectivity struct {
	Iface   string `json:"iface,omitempty"`
	LocalIP string `json:"local_ip,omitempty"`
	MAC     string `json:"mac,omitempty"`
}

func (c Connectivity) MarshalJSON() ([]byte, error) {
	if c == (Connectivity{}) {
		return []byte("{}"), nil
	}
	type ConnectivityAlias Connectivity
	return json.Marshal(ConnectivityAlias(c))
}

type Location struct {
	Longitude string     `json:"longitude"`
	Latitude  string     `json:"latitude"`
	Altitude  string     `json:"altitude"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Action struct {
	ServiceAction string `json:"service_action"`
}

type Sensors struct {
	Name string
}

type InjectionPlugin struct {
	Plugin    string     `json:"plugin"`
	Status    string     `json:"status"`
	OnBoot    bool       `json:"onboot"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
