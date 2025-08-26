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

// UpdateServer updates an existing server record in the database.
func UpdateServer(server *Server) error {
	db := GetDB()
	return db.Save(server).Error
}

// DeleteServer deletes a server record by ID.
func DeleteServer(id uint) error {
	db := GetDB()
	return db.Delete(&Server{}, id).Error
}
