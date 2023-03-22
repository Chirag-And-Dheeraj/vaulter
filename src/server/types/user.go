package types

type User struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" validate:"required,email"`
	Pin          string `json:"pin" validate:"required,min=6,max=6"`
	ProfilePhoto string `json:"profile_photo"`
}
