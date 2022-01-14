package system

import (
	"dogrod-web-service/global"
	"dogrod-web-service/model/system"
	"dogrod-web-service/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SystemUserService struct{}

// sign up user
func (systemUserService *SystemUserService) SignUp(u system.SystemUser) (err error, userInter system.SystemUser) {
	var existUser system.SystemUser

	if !errors.Is(global.ConnectDatabase().Where("username = ?", u.Username).First(&existUser).Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("Username already exist"), userInter
	}
	u.Password = utils.MD5V([]byte(u.Password))

	err = global.ConnectDatabase().Create(&u).Error
	return err, u
}
