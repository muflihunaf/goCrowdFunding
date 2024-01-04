package user

type UserFormatter struct {
	ID         int
	Name       string
	Occupation string
	Email      string
	Token      string
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formatter
}
