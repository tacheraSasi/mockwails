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