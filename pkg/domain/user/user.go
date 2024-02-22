package user

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Status   bool      `json:"status"`
}

func (u *User) GetId() uuid.UUID {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetStatus() bool {
	return u.Status
}

func NewUser(email string, password string) (User, error) {
	if email == "" {
		return User{}, errors.New("empty email")
	}

	return User{
		Id:       uuid.New(),
		Email:    email,
		Password: password,
		Status:   true,
	}, nil
}
