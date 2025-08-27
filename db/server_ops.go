package db

// CreateServer inserts a new server record into the database.
func CreateServer(server *Server) error {
	db := GetDB()
	return db.Create(server).Error
}

// GetAllServers retrieves all server records from the database.
func GetAllServers() ([]Server, error) {
	db := GetDB()
	var servers []Server
	err := db.Find(&servers).Error
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
