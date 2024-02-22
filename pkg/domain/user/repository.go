package user

import "github.com/google/uuid"

type UserRepository interface {
	AddUser(User) error
	GetUser(uuid.UUID) (User, error)
}
