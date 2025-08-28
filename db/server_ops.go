package db



// CreateServer inserts a new server record into the database.
func CreateServer(server *Server) error {
	db := GetDB()

	var existingAddress AddressAssigned
	err := db.Where("port = ?", server.AddressAssigned.Port).First(&existingAddress).Error
	if err == nil {
		// Port is already assigned, find next available port
		server.AddressAssigned.Port = findNextAvailablePort(server.AddressAssigned.Port)
	}

	return db.Create(server).Error
}

// findNextAvailablePort finds the next available port starting from the given port
func findNextAvailablePort(startPort int) int {
	//TODO:: Maybe i will free all the running port that are not in use and started by mockwails
	db := GetDB()
	for port := startPort; port <= 65535; port++ {
		var existingAddress AddressAssigned
		err := db.Where("port = ?", port).First(&existingAddress).Error
		if err != nil { // Port not found in database, it's available
			return port
		}
	}
	// If no port is available in the range, start from 8000
	for port := 8000; port < startPort; port++ {
		var existingAddress AddressAssigned
		err := db.Where("port = ?", port).First(&existingAddress).Error
		if err != nil {
			return port
		}
	}
	return 9000 // fallback port
}

// Check if endpoint exists
func CheckIfEndpointExists(endpoint string, port int) (bool, error) {
	db := GetDB()
	var count int64
	err := db.Model(&Server{}).Where("endpoint = ? AND port = ?", endpoint, port).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetAllServers retrieves all server records from the database.
func GetAllServers() ([]Server, error) {
	db := GetDB()
	var servers []Server
	err := db.Find(&servers).Error
	return servers, err
}

// GetAllActiveServers retrieves all active server records from the database.
func GetAllActiveServers() ([]Server, error) {
	db := GetDB()
	var servers []Server
	err := db.Where("status = ?", "active").Find(&servers).Error
	return servers, err
}

func GetServerByID(id uint) (*Server, error) {
	db := GetDB()
	var server Server
	err := db.Preload("AddressAssigned").First(&server, id).Error
	if err != nil {
		return nil, err
	}
	return &server, nil
}

// UpdateServer updates an existing server record in the database.
func UpdateServer(server *Server) error {
	db := GetDB()
	return db.Save(server).Error
}

// ToggleServerStatus toggles the status of a server between active and inactive.
func ToggleServerStatus(id uint) error {
	db := GetDB()
	var server Server
	if err := db.First(&server, id).Error; err != nil {
		return err
	}
	if server.Status == "active" {
		server.Status = "inactive"
	} else {
		server.Status = "active"
	}
	return db.Save(&server).Error
}

// DeleteServer deletes a server record by ID.
func DeleteServer(id uint) error {
	db := GetDB()
	return db.Delete(&Server{}, id).Error
}
