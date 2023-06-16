package connection

import (
	"fly/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	dsn := "root:NvQaXmLx7pTrABi4iSC7@tcp(containers-us-west-8.railway.app:7502)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if err := db.AutoMigrate(&user.Users{}); err != nil {
		return db, err
	}

	return db, nil
}
