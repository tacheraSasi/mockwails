package db

// SeedServers creates initial seed servers if they do not already exist (no duplicates by Name).
func SeedServers() error {
	db := GetDB()

	// Fix existing servers that don't have AddressAssigned
	var serversWithoutAddress []Server
	err := db.Where("id NOT IN (SELECT server_id FROM address_assigned)").Find(&serversWithoutAddress).Error
	if err == nil {
		for _, server := range serversWithoutAddress {
			port := findNextAvailablePort(8000)
			addressAssigned := AddressAssigned{
				ServerID: server.ID,
				Port:     port,
			}
			db.Create(&addressAssigned)
		}
	}

	seedData := []Server{
		{
			Name:            "Sample User API",
			Endpoint:        "/api/users",
			Method:          "GET",
			Description:     "Returns a list of users",
			ResponseStatus:  200,
			Status:          "inactive",
			ResponseBody:    `[{"id":1,"name":"John Doe"}]`,
			RequestHeaders:  "",
			RequestQuery:    "",
			ResponseHeaders: "",
			RequestBody:     "",
			AddressAssigned: AddressAssigned{Port: 8001},
		},
		{
			Name:            "Sample Product API",
			Endpoint:        "/api/products",
			Method:          "POST",
			Description:     "Creates a new product",
			ResponseStatus:  201,
			Status:          "inactive",
			RequestBody:     `{"name":"New Product"}`,
			ResponseBody:    `{"id":100,"name":"New Product"}`,
			RequestHeaders:  "",
			RequestQuery:    "",
			ResponseHeaders: "",
			AddressAssigned: AddressAssigned{Port: 8002},
		},
		{
			Name:            "Update User API",
			Endpoint:        "/api/users/{id}",
			Method:          "PUT",
			Description:     "Updates a user by ID",
			ResponseStatus:  200,
			Status:          "inactive",
			RequestBody:     `{"name":"Jane Doe"}`,
			ResponseBody:    `{"id":1,"name":"Jane Doe"}`,
			RequestHeaders:  "",
			RequestQuery:    "",
			ResponseHeaders: "",
			AddressAssigned: AddressAssigned{Port: 8003},
		},
		{
			Name:            "Delete Product API",
			Endpoint:        "/api/products/{id}",
			Method:          "DELETE",
			Description:     "Deletes a product by ID",
			ResponseStatus:  204,
			Status:          "inactive",
			RequestBody:     "",
			ResponseBody:    "",
			RequestHeaders:  "",
			RequestQuery:    "",
			ResponseHeaders: "",
			AddressAssigned: AddressAssigned{Port: 8004},
		},
		{
			Name:            "Update User Email API",
			Endpoint:        "/api/users/{id}/email",
			Method:          "PATCH",
			Description:     "Updates a user's email address",
			ResponseStatus:  200,
			Status:          "inactive",
			RequestBody:     `{"email":"new@email.com"}`,
			ResponseBody:    `{"id":1,"email":"new@email.com"}`,
			RequestHeaders:  "",
			RequestQuery:    "",
			ResponseHeaders: "",
			AddressAssigned: AddressAssigned{Port: 8005},
		},
		{
			Name:            "Get Orders API",
			Endpoint:        "/api/orders",
			Method:          "GET",
			Description:     "Returns a list of orders",
			ResponseStatus:  200,
			Status:          "inactive",
			RequestBody:     "",
			ResponseBody:    `[{"id":1,"total":99.99}]`,
			RequestHeaders:  "",
			RequestQuery:    "",
			ResponseHeaders: "",
			AddressAssigned: AddressAssigned{Port: 8006},
		},
	}

	for _, s := range seedData {
		var count int64
		if err := db.Model(&Server{}).Where("name = ?", s.Name).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			if err := db.Create(&s).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

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
