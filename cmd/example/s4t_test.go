package example

import (
	"fmt"
	"s4t-sdk-module/pkg"
	"s4t-sdk-module/pkg/api/boards"
	"s4t-sdk-module/pkg/api/plugins"
	"s4t-sdk-module/pkg/api/services"
	"testing"

)

var service_id = ""
var board_id = "c910e7f1-74d0-4f76-ae6a-a46c1da0d92d"
var plugin_id = ""

func TestGetBoards(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	board := boards.Board{}
	resp, err := board.ListBoards(client)
	
	if err != nil {
		t.Errorf("Error listing boards: %v", err)
	}

	for _, board := range resp {
		fmt.Printf("Board Name: %s, Status: %s\n", board.Name, board.Status)
	}
}

func TestGetBoardDetails(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	

	board := boards.Board{}
	resp, err := board.GetBoardDetail(client, "c910e7f1-74d0-4f76-ae6a-a46c1da0d92d")
	
	if err != nil {
		t.Errorf("Error getting board info: %v", err)
	}

	fmt.Printf("Board Name: %s, Status: %s\n", resp.Name, resp.Status)

}

/*
func TestCreateBoard(t *testing.T) {
	auth_req, err := read_config.ReadConfiguration()
	
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	client := s4t.NewClient("http://" + auth_req.S4tAuthData.Ip)
	client.AuthToken = auth_req.S4tAuthData.Token	
	
	test_board := boards.Board{
		Uuid: "6ba7b810-9dad-11d1-80b4-00c04fd430c9",
		Code: "demo-test-generic",
		Status: "offline",
		Name: "s4t-sdk-testing-board",
		Type: "gateway",
		Agent: "wagent1",
		// Wstunip: "172.17.4.2",
		Session: "7712408803711087",
		Fleet: nil,
		LRversion: "0.4.17",
		Connectivity: boards.Connectivity{},
		Location: []boards.Location{
			{ 
			Longitude: "1.0",
			Latitude: "1.0",
			Altitude: "1.0",
		}},
	}

	_, err = test_board.CreateBoard(client, test_board)
	
	if err != nil {
		t.Errorf("Error creating board: %v", err)
	}

}
*/

func TestPatchBoard(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	

	board := boards.Board{}
	updated_board_data := map[string]interface{}{
        "code": "test-generic-patched",
    }

	resp, err := board.PatchBoard(client, "6ba7b810-9dad-11d1-80b4-00c04fd430c9",updated_board_data)
	
	if err != nil {
		t.Errorf("Error creating board: %v", err)
	}
	
	fmt.Printf("Board Name: %s, Status: %s\n", resp.Name, resp.Code)
} 

/*
// REQUIRE THE CORRECT ACTION IF NOT RETURN ERROR
func TestBoardAction(t *testing.T) {
	client, err := s4t.GetClientConnection()

	board := boards.Board{}
	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	
	action_data := map[string]interface{}{
        "action": "test-action",
		"parameters": map[string] interface{} {},
	}

	err = boards.PerformBoardAction(client, "6ba7b810-9dad-11d1-80b4-00c04fd430c9", action_data)
	
	if err != nil {
		t.Errorf("Error creating board: %v", err)
	}

}
*/

func TestGetServices(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	
	service := services.Service{}
	resp, err := service.GetServices(client)
	
	if err != nil {
		t.Errorf("Error getting service info: %v", err)
	}

	for _, service := range resp {
		fmt.Printf("Service Name: %s, Status: %s\n", service.Name, service.Uuid)
	}
}

func TestCreateService(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	
	service := services.Service {
		Name: "test_service",
		Port: 4321,
		Protocol: "TCP",

	}

	resp, err := service.CreateService(client, service)
	
	if err != nil {
		t.Errorf("Error creating service info: %v", err)
	}

	fmt.Printf("Service Name: %s, Status: %s\n", resp.Name, resp.Uuid)

	service_id = resp.Uuid
}

func TestPatchService(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	
	updated_service_data := map[string]interface{}{
		"name": "test-service-generic-patched",
	}

	service := services.Service{}
	resp, err := service.PatchService(client, service_id, updated_service_data)
	
	if err != nil {
		t.Errorf("Error creating service info: %v", err)
	}

	fmt.Printf("Service Name: %s\n", resp.Name)
}

func TestDeleteService(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	
	service := services.Service{}
	err = service.DeleteService(client, service_id)
	
	if err != nil {
		t.Errorf("Error creating service info: %v", err)
	}

}

func TestBoardExposedServices(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	

	service := services.Service{}
	resp, err := service.GetBoardExposedServices(client, board_id)
	
	if err != nil {
		t.Errorf("Error getting service info: %v", err)
	}

	for _, service := range resp {
		fmt.Printf("Service Name: %s, Status: %s\n", service.Name, service.Uuid)
	}
}

func TestRestoreBoardService(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	

	service := services.Service{}
	err = service.RestoreService(client, board_id)
	
	if err != nil {
		t.Errorf("Error getting service info: %v", err)
	}
}

/*
func TestPerformActionOnService(t *testing.T) {
	client, err := s4t.GetClientConnection()

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
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	plugin := plugins.Plugin{}
	resp, err := plugin.GetPlugins(client)
	
	if err != nil {
		t.Errorf("Error getting plugin info: %v", err)
	}

	for _, plugin := range resp {
		fmt.Printf("Plugin Name: %s, Status: %s\n", plugin.Name, plugin.UUID)
	}

}

func TestGetPlugin(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	

	plugin := plugins.Plugin{}
	resp, err := plugin.GetPlugin(client, "b5217ab0-82e9-46c0-94d6-1c0d79437db6")
	
	if err != nil {
		t.Errorf("Error getting plugin info: %v", err)
	}

	fmt.Printf("Plugin Name: %s, Status: %s\n", resp.Name, resp.UUID)
}

// CANNOT LOAD CODE IN THE BASE CLASS
func TestCreatePlugin(t *testing.T) {	
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	
	plugin_req := plugins.PluginReq{
		Name: "Test-plugin-s4t",
		Parameters: map[string]interface{}{},
		Code:"from iotronic_lightningrod.plugins import Plugin\n\nfrom oslo_log import log as logging\n\nLOG = logging.getLogger(__name__)\n\n\n# User imports\n\n\nclass Worker(Plugin.Plugin):\n    def __init__(self, uuid, name, q_result, params=None):\n        super(Worker, self).__init__(uuid, name, q_result, params)\n\n    def run(self):\n        LOG.info(\"Input parameters: \" + str(self.params))\n        LOG.info(\"Plugin \" + self.name + \" process completed!\")\n        self.q_result.put(\"ZERO RESULT\")",
		// Description: "A generic test plugin",
	}
	
	plugin := plugins.Plugin{}
	resp, err := plugin.CreatePlugin(client, plugin_req)
	
	if err != nil {
		t.Errorf("Error creating plugin: %v", err)
	}
	
	plugin_id = resp.UUID
	fmt.Printf("Plugin name: %v", resp.Name)
}

func TestPatchPlugin(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	
	updated_service_data := map[string]interface{}{
		"name": "test-plugin-generic-patched",
	}

	plugin := plugins.Plugin{}
	resp, err := plugin.PacthPlugin(client, plugin_id, updated_service_data)
	
	if err != nil {
		t.Errorf("Error patching plugin info: %v", err)
	}

	fmt.Printf("Plugin Name: %s\n", resp.Name)
}

func TestInjectBoardPlugin(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	

	data := map[string] interface{} {
		"plugin": plugin_id,
		// "onboot": "yes",
		// "force": "yes",

	} 

	plugin := plugins.Plugin{}
	err = plugin.InjectPLuginBoard(client, board_id, data)
	
	if err != nil {
		t.Errorf("Error getting plugin info: %v", err)
	}
}

func TestDeleteBoardPlugin(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	
	plugin := plugins.Plugin{}
	err = plugin.RemoveInjectedPlugin(client, plugin_id, board_id)
	
	if err != nil {
		t.Errorf("Error deleting plugin: %v", err)
	}
}


func TestDeletePLugin(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	
	
	plugin := plugins.Plugin{}
	err = plugin.DeletePlugin(client, plugin_id)
	
	if err != nil {
		t.Errorf("Error deleting plugin: %v", err)
	}

}

func TestGetBoardPlugins(t *testing.T) {
	c := s4t.Client{}
	client, err := c.GetClientConnection()

	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}	

	plugin := plugins.Plugin{}
	resp, err := plugin.GetBoardPlugins(client, board_id)
	
	if err != nil {
		t.Errorf("Error getting plugin info: %v", err)
	}

	for _, plugin := range resp {
		fmt.Printf("Plugin Name: %s, Status: %s\n", plugin.Name, plugin.UUID)
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

