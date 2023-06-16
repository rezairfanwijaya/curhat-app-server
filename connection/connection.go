package connection

import (
	"fly/user"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(envPath string) (*gorm.DB, error) {
	env, err := godotenv.Read(envPath)
	if err != nil {
		return &gorm.DB{}, err
	}

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		env["DATABASE_USERNAME"],
		env["DATABASE_PASSWORD"],
		env["DATABASE_HOST"],
		env["DATABASE_PORT"],
		env["DATABASE_NAME"],
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if err := db.AutoMigrate(&user.Users{}); err != nil {
		return db, err
	}

	if err := MigrationUser(db); err != nil {
		return db, err
	}

	return db, nil
}
