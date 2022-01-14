package system

import (
	"database/sql"
	"dogrod-web-service/global"
)

type SystemUser struct {
	global.BaseModel
	Username    string       `json:"username" gorm:"comment:unique user name"`
	Email       string       `json:"email" gorm:"email"`
	Password    string       `json:"password" gorm:"password"`
	LastLogin   sql.NullTime `json:"lastLogin" gorm:"user last login time"`
	IsSuperuser bool         `json:"isSuperuser" gorm:"is super user"`
	IsActive    bool         `json:"isActive" gorm:"is user active"`
	IsAdminUser bool         `json:"isAdminUser" gorm:"is user have admin permission"`
}
