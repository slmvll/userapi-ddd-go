package user

type UserRepository interface {
	AddUser(User) error
}
