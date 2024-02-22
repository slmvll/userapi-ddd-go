package user

import "github.com/google/uuid"

type UserService interface {
	AddUser(User) (uuid.UUID, error)
}

type userService struct {
	r UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{r}
}

func (us *userService) AddUser(u User) (uuid.UUID, error) {

	newUser, err := NewUser(u.Email, u.Password)
	if err != nil {
		return uuid.UUID{}, err
	}
	// TO-DO: If password is empty -> autogenerate
	if errAdd := us.r.AddUser(newUser); errAdd != nil {
		return uuid.UUID{}, errAdd
	}

	return newUser.Id, nil
}
