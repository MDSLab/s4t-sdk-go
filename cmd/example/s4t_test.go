package example

import (
	"fmt"
	"math/rand"
	"s4t-sdk-module/pkg"
	"s4t-sdk-module/pkg/api/boards"
	"s4t-sdk-module/pkg/read_conf"
	"testing"
	"time"
)

func randStringGenerator(n int) string {
    
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	
	for i := range b {
        b[i] = letters[seededRand.Intn(len(letters))]
    }
    return string(b)
}


func TestGetBoards(t *testing.T) {
	auth_req, err := read_config.ReadConfiguration()
	
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	client := s4t.NewClient("http://" + auth_req.S4tAuthData.Ip)
	client.AuthToken = auth_req.S4tAuthData.Token	

	resp, err := boards.ListBoards(client)
	
	if err != nil {
		t.Errorf("Error listing boards: %v", err)
	}

	for _, board := range resp {
		fmt.Printf("Board Name: %s, Status: %s\n", board.Name, board.Status)
	}
}

func TestGetBoardDetails(t *testing.T) {
	auth_req, err := read_config.ReadConfiguration()
	
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	client := s4t.NewClient("http://" + auth_req.S4tAuthData.Ip)
	client.AuthToken = auth_req.S4tAuthData.Token	

	resp, err := boards.GetBoardDetail(client, "c910e7f1-74d0-4f76-ae6a-a46c1da0d92d")
	
	if err != nil {
		t.Errorf("Error getting board info: %v", err)
	}

	fmt.Printf("Board Name: %s, Status: %s\n", resp.Name, resp.Status)

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

	_, err = boards.CreateBoard(client, test_board)
	
	if err != nil {
		t.Errorf("Error creating board: %v", err)
	}

}
*/

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


