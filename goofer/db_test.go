package goofer

import (
	"os"
	"testing"
)

func TestServerCRUD(t *testing.T) {
	// Clean up test db before and after
	os.Remove(DB_PATH)
	defer os.Remove(DB_PATH)

	server := ServerEntity{
		Name:            "TestServer",
		Description:     "A test server",
		Endpoint:        "/api/test",
		Method:          "GET",
		RequestHeaders:  "{}",
		RequestBody:     "{}",
		ResponseStatus:  200,
		ResponseHeaders: "{\"Content-Type\": \"application/json\"}",
		ResponseBody:    "{\"message\": \"ok\"}",
	}

	// Create
	err := CreateServer(server)
	if err != nil {
		t.Fatalf("CreateServer failed: %v", err)
	}

	// Read
	servers, err := GetAllServers()
	if err != nil {
		t.Fatalf("GetAllServers failed: %v", err)
	}
	if len(servers) == 0 {
		t.Fatal("No servers found after creation")
	}
	found := servers[0]
	if found.Name != server.Name {
		t.Errorf("Expected Name %q, got %q", server.Name, found.Name)
	}

	// Update
	found.Description = "Updated description"
	err = UpdateServer(found)
	if err != nil {
		t.Fatalf("UpdateServer failed: %v", err)
	}
	servers, _ = GetAllServers()
	if servers[0].Description != "Updated description" {
		t.Errorf("Update did not persist")
	}

	// Delete
	err = DeleteServer(found.ID)
	if err != nil {
		t.Fatalf("DeleteServer failed: %v", err)
	}
	servers, _ = GetAllServers()
	if len(servers) != 0 {
		t.Errorf("Server not deleted")
	}
}
