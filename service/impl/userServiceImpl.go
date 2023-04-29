package impl

import (
	"newim/dao"
	"newim/service"
)

type UserServiceImpl struct {
}

var _ service.UserService = (*UserServiceImpl)(nil)

func (u UserServiceImpl) UserRegister(userId, pwd string) error {
	return dao.UserRegister(userId, pwd)
}
