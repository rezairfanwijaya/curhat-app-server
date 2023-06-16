package user

type userFormatted struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func FormatterUser(user User) *userFormatted {
	return &userFormatted{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
}

func FormatterUsers(users []User) []*userFormatted {
	var result []*userFormatted

	for _, user := range users {
		userFormatted := FormatterUser(user)
		result = append(result, userFormatted)
	}

	return result
}
