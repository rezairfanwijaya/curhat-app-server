package user

func SeederUser() *[]Users {
	return &[]Users{
		{
			ID:       1,
			Email:    "user1@gmail.com",
			Role:     "user",
			Password: "12345",
		}, {
			ID:       2,
			Email:    "user2@gmail.com",
			Role:     "user",
			Password: "12345",
		}, {
			ID:       3,
			Email:    "admin@admin.com",
			Role:     "admin",
			Password: "12345",
		},
	}
}
