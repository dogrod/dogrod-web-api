package user

import (
	"dogrod-web-service/global"
	"dogrod-web-service/model/user"
	"dogrod-web-service/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserBaseService struct{}

// sign up user
func (userBaseService *UserBaseService) SignUp(u user.UserBase) (err error, internalUser user.UserBase) {
	var existUser user.UserBase

	if !errors.Is(global.ConnectDatabase().Where("username = ?", u.Username).First(&existUser).Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("Username already exist"), internalUser
	}
	u.Password = utils.MD5V([]byte(u.Password))

	err = global.ConnectDatabase().Create(&u).Error
	return err, u
}
