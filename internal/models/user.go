package models

type User struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string  `json:"last_name" validate:"required,min=2,max=50"`
	Email     string  `json:"email" validate:"required,email"`
	Balance   float64 `json:"balance"`
}
