package request

import "flightbooking/db"

type RegisterUserRequest struct {
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password" validate:"min=3"`
	FirstName string    `json:"firstName" validate:"min=3"`
	LastName  string    `json:"lastName" validate:"min=3"`
	Gender    db.Gender `json:"gender"`
}

type UpdateUserRequest struct {
	FirstName string    `json:"firstName" validate:"required;min=3"`
	LastName  string    `json:"lastName" validate:"required;min=3"`
	Gender    db.Gender `json:"gender"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ValidatePasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
