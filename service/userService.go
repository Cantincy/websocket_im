package service

type UserService interface {
	UserRegister(userId, pwd string) error
}
