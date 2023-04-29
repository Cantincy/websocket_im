package dao

import (
	"fmt"
	"newim/entity"
)

func UserRegister(userId, pwd string) error {
	var count int64
	err := DB.Debug().Model(&entity.User{}).Where("userId=? and pwd=?", userId, pwd).Count(&count).Error
	if err != nil {
		return err
	}
	if count != 0 {
		return fmt.Errorf("userId already exist")
	}
	err = DB.Debug().Model(&entity.User{}).Create(&entity.User{UserId: userId, Pwd: pwd}).Error
	return err
}
