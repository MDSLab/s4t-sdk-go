package example

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MIKE9708/s4t-sdk-go/pkg/api"
	boards "github.com/MIKE9708/s4t-sdk-go/pkg/api/data/board"
	"github.com/MIKE9708/s4t-sdk-go/pkg/api/data/plugin"
	"github.com/MIKE9708/s4t-sdk-go/pkg/api/data/service"
	read_config "github.com/MIKE9708/s4t-sdk-go/pkg/read_conf"
	"k8s.io/apimachinery/pkg/runtime"
)

var service_id = ""
var board_data = ""
var plugin_data = ""
var f interface{}

func TestGetBoards(t *testing.T) {
	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	resp, err := client.ListBoards()

	if err != nil {
		t.Errorf("Error listing boards: %v", err)
	}

	for _, board := range resp {
		fmt.Printf("Test Get board returned UUID: %s  Name: %s, Status: %s\n\n", board.Uuid, board.Name, board.Status)
	}
}

func TestGetBoardDetails(t *testing.T) {
	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	board := "361be02d-3324-4746-a8a0-0ee049ea1dad"
	resp, err := client.GetBoardDetail(board)
	board_data = board
	if err != nil {
		t.Errorf("Error getting board info: %v", err)
	}

	fmt.Printf("Board Name: %s, Status: %s\n\n", resp.Name, resp.Status)
}

func TestCreateBoard(t *testing.T) {
	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	test_board := boards.Board{
		Code: "erjfiorejfiojre",
		Name: "jernorefo",
		Type: "gateway",
		Location: []*boards.Location{
			{
				Longitude: "1.0",
				Latitude:  "1.0",
				Altitude:  "1.0",
			},
		},
	}

	// board_data := map[string]interface{}{
	// 	"code": "demo-test-generic",
	// 	"name": "s4t-demo",
	// 	"type": "gateway",
	// 	"location": []map[string]string{{
	// 		"longitude": "1.0",
	// 		"latitude":  "1.0",
	// 		"altitude":  "1.0",
	// 	}},
	// }

	_, err = client.CreateBoard(test_board)

	if err != nil {
		t.Errorf("Error creating board: %v", err)
	}

}

/*
func TestPatchBoard(t *testing.T) {
	c := s4t.Client{}


	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	updated_board_data := map[string]interface{}{
		"code": "test-patched",
	}

	resp, err := client.PatchBoard(board_data, updated_board_data)

	if err != nil {
		t.Errorf("Error patching board: %v", err)
	}

	fmt.Printf("Board Name: %s, Status: %s\n\n", resp.Name, resp.Code)
}
*/

// REQUIRE THE CORRECT ACTION IF NOT RETURN ERROR
func TestBoardAction(t *testing.T) {
	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	action_data := map[string]interface{}{
		"action": "ServiceEnable",
	}
	err = client.PerformBoardAction("6ba7b810-9dad-11d1-80b4-00c04fd430c9", service_id, action_data)

	if err != nil {
		t.Errorf("Error creating board: %v", err)
	}
}

func TestGetService(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	service := "be86610b-d401-416b-ac11-7c1183019830"
	resp, err := client.GetService(service)

	if err != nil {
		t.Errorf("Error getting plugin info: %v", err)
	}

	t.Logf("Plugin Name: %s, Status: %s\n\n", resp.Name, resp.Uuid)
}

func TestGetServices(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	resp, err := client.GetServices()

	if err != nil {
		t.Errorf("Error getting service info: %v", err)
	}

	for _, service := range resp {
		fmt.Printf("Service Name: %s, Status: %s\n\n", service.Name, service.Uuid)
	}
}

func TestCreateService(t *testing.T) {
	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	service := services.Service{
		Name:     "test_service",
		Port:     4321,
		Protocol: "TCP",
	}

	resp, err := client.CreateService(service)

	if err != nil {
		t.Errorf("Error creating service info: %v", err)
	}

	fmt.Printf("Service Name: %s, Status: %s\n\n", resp.Name, resp.Uuid)

	service_id = resp.Uuid
}

func TestPatchService(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	updated_service_data := map[string]interface{}{
		"name": "test-service-generic-patched",
	}

	resp, err := client.PatchService(service_id, updated_service_data)

	if err != nil {
		t.Errorf("Error creating service info: %v", err)
	}

	fmt.Printf("Service Name: %s\n\n", resp.Name)
}

func TestDeleteService(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	err = client.DeleteService(service_id)

	if err != nil {
		t.Errorf("Error creating service info: %v", err)
	}

}

func TestBoardExposedServices(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	resp, err := client.GetBoardExposedServices(board_data)

	if err != nil {
		t.Errorf("Error getting service info: %v", err)
	}

	for _, service := range resp {
		fmt.Printf("Service Name: %s, Status: %s\n\n", service.Name, service.Uuid)
	}
}

func TestRestoreBoardService(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	err = client.RestoreService(board_data)

	if err != nil {
		t.Errorf("Error getting service info: %v", err)
	}
}

/*
func TestPerformActionOnService(t *testing.T) {
	client, err := s4t.s4t.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	action := boards.Action {
		ServiceAction: "test",
	}

	err = services.PerfomActionOnService(client, board_id, service_id, action)

	if err != nil {
		t.Errorf("Error getting service info: %v", err)
	}


}
*/

func TestGetPlugins(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	resp, err := client.GetPlugins()

	if err != nil {
		t.Errorf("Error getting plugin info: %v", err)
	}

	for _, plugin := range resp {
		t.Logf("Plugin Name: %s, Status:%s\n\n", plugin.Name, plugin.UUID)
	}

}

func TestGetPlugin(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	plugin := "5a0b6644-7442-4d6c-9b88-4577f14faea6"
	resp, err := client.GetPlugin(plugin)

	if err != nil {
		t.Errorf("Error getting plugin info: %v", err)
	}

	t.Logf("Plugin Name: %s, Status: %s\n\n", resp.Name, resp.UUID)
}

// CANNOT LOAD CODE IN THE BASE CLASS
func TestCreatePlugin(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	err = json.Unmarshal([]byte(`{}`), &f)

	plugin_req := plugins.PluginReq{
		Name:       "Test-plugin-s4t",
		Parameters: runtime.RawExtension{Raw: []byte(`{}`)},
		Code:       "from iotronic_lightningrod.plugins import Plugin\n\nfrom oslo_log import log as logging\n\nLOG = logging.getLogger(__name__)\n\n\n# User imports\n\n\nclass Worker(Plugin.Plugin):\n    def __init__(self, uuid, name, q_result, params=None):\n        super(Worker, self).__init__(uuid, name, q_result, params)\n\n    def run(self):\n        LOG.info(\"Input parameters: \" + str(self.params))\n        LOG.info(\"Plugin \" + self.name + \" process completed!\")\n        self.q_result.put(\"ZERO RESULT\")",
		// Description: "A generic test plugin",
	}

	resp, err := client.CreatePlugin(plugin_req)

	if err != nil {
		t.Errorf("Error creating plugin: %v", err)
	}

	plugin_data = resp.UUID
	t.Logf("Plugin name: %v\n\n", resp.Name)
}

func TestPatchPlugin(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	updated_service_data := map[string]interface{}{
		"name": "test-plugin-generic-patched",
	}

	resp, err := client.PacthPlugin(plugin_data, updated_service_data)

	if err != nil {
		t.Errorf("Error patching plugin info: %v", err)
	}

	t.Logf("Plugin Name: %s\n\n", resp.Name)
}

func TestInjectBoardPlugin(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	data := map[string]interface{}{
		// "plugin": plugin_data,
		"uuid":       "5a0b6644-7442-4d6c-9b88-4577f14faea6",
		"board_list": "e2b967e7-f643-4840-b67e-489524cb2ee5",
		"name":       "test-plugin-generic-patched",
		// "onboot": "yes",
		// "force": "yes",
	}

	err = client.InjectPLuginBoard(board_data, data)

	if err != nil {
		t.Errorf("Error getting plugin info: %v", err)
	}
}

func TestGetBoardPlugins(t *testing.T) {
	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	resp, err := client.GetBoardPlugins(board_data)

	if err != nil {
		t.Errorf("Error getting plugin info: %v", err)
	}

	for _, plugin := range resp {
		t.Logf("\n\n\nBoard Plugin UUID: %s\n\n", plugin.Plugin)
	}
}
func TestDeleteBoardPlugin(t *testing.T) {

	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	err = client.RemoveInjectedPlugin(plugin_data, board_data)

	if err != nil {
		t.Errorf("Error deleting plugin: %v", err)
	}
}

func TestDeletePLugin(t *testing.T) {
	authreq := read_config.FormatAuthRequ(
		"admin",
		"ADMIN_PASS",
		"default",
	)
	client, err := s4t.GetClientConnection(*authreq)

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	err = client.DeletePlugin(plugin_data)

	if err != nil {
		t.Errorf("Error deleting plugin: %v", err)
	}

}

/*
// ERROR WHEN CALLING DELETE "catching classes that do not inherit from BaseException is not allowed\"
func TestDeleteBoard(t *testing.T) {
	auth_req, err := read_config.ReadConfiguration()

	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	client := s4t.NewClient("http://" + auth_req.S4tAuthData.Ip)
	client.AuthToken = auth_req.S4tAuthData.Token

	err = boards.DeleteBoard(client, "6ba7b810-9dad-11d1-80b4-00c04fd430c8")

	if err != nil {
		t.Errorf("Error deleting board: %v", err)
	}

}

// 404 NOT FOUND
func TestAddNewPosition(t *testing.T) {
	auth_req, err := read_config.ReadConfiguration()

	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	client := s4t.NewClient("http://" + auth_req.S4tAuthData.Ip)
	client.AuthToken = auth_req.S4tAuthData.Token

	position := boards.Location{
		Altitude: "2.0",
		Latitude: "1.0",
		Longitude: "1.0",
	}
	err = boards.AddNewPosition(client, "c910e7f1-74d0-4f76-ae6a-a46c1da0d92d", position)

	if err != nil {
		t.Errorf("Error getting board info: %v", err)
	}

}
*/
