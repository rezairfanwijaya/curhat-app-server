package connection

import (
	"fly/user"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(envPath string) (*gorm.DB, error) {
	env, _ := godotenv.Read(envPath)

	username := env["DATABASE_USERNAME"]
	password := env["DATABASE_PASSWORD"]
	host := env["DATABASE_HOST"]
	port := env["DATABASE_PORT"]
	databaseName := env["DATABASE_NAME"]

	if username == "" {
		os.Getenv("DATABASE_USERNAME")
	}
	if password == "" {
		os.Getenv("DATABASE_PASSWORD")
	}
	if host == "" {
		os.Getenv("DATABASE_HOST")
	}
	if port == "" {
		os.Getenv("DATABASE_PORT")
	}
	if databaseName == "" {
		os.Getenv("DATABASE_NAME")
	}

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		databaseName,
	)

	log.Println(dsn)

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
