package user

import "time"

type CreateUserRequestDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserResponseDTO struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserCommand struct {
	Email string
	Name  string
}
