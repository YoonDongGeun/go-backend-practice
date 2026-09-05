package user

import "time"

type UserEntity struct {
	ID        int64
	Email     string
	Name      string
	CreatedAt time.Time
}
