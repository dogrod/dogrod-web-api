package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "dev_dogrod_com:fut@rfe4FWT_fuy0jun@tcp(sh-cdb-lldgzazy.sql.tencentcdb.com:59717)/dev_dogrod_com"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
