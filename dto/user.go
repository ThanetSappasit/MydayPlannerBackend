package dto

type GetUserByEmail struct {
	Email *string `json:"email"`
}

type AccUserRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	HashedPassword string `json:"hashed_password" binding:"required"`
	Profile        string `json:"profile"`
}

type DeleteUserRequest struct {
	Email string `json:"email"`
}
