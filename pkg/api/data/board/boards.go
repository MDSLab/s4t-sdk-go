package boards
import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/MIKE9708/s4t-sdk-go/pkg/api/data/generic"
)
// +kubebuilder:object:generate=true
type Board struct {
	Uuid string `json:"uuid"`
	Code string `json:"code"`
	Status string `json:"status"`
	Name string `json:"name"`
	Type string `json:"type"`
	Agent string `json:"agent"`
	Wstunip string `json:"wstun_ip,omitempty"`
	Session string `json:"session"`
	Fleet runtime.RawExtension `json:"fleet,omitempty"`
	//interface{} `json:"fleet"`
	LRversion string `json:"lr_version"`
	Connectivity Connectivity `json:"connectivity"`
	Links []generic.Link `json:"links,omitempty"`
	Location []Location  `json:"location"`
}
func (b *Board) Keys() []string {
    return  []string{
		"uuid", "code", 
		"status", "name", 
		"type", "agent", 
		"wstpun_ip","session",
		"fleet","lr_version",
		"connectivity","links",
		"location",
	}
}

type Connectivity struct {
    Iface   string `json:"iface"`
    LocalIP string `json:"local_ip"`
    MAC     string `json:"mac"`
}
func (c Connectivity) MarshalJSON() ([]byte, error) {
	if c == (Connectivity{}) {
		return []byte("{}"), nil
	}
	type ConnectivityAlias Connectivity
	return json.Marshal(ConnectivityAlias(c))
}




type Location struct {
	Longitude  string      `json:"longitude"`
	Latitude   string      `json:"latitude"`
	Altitude   string      `json:"altitude"`
	UpdatedAt  []byte  `json:"updated_at"`
}

type Action struct {
	ServiceAction string `json:"service_action"`
}

type Sensors struct {
	Name string
}


