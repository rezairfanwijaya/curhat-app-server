package connection

import (
	"fly/note"
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
		username = os.Getenv("MYSQLUSER")
	}
	if password == "" {
		password = os.Getenv("MYSQLPASSWORD")
	}
	if host == "" {
		host = os.Getenv("MYSQLHOST")
	}
	if port == "" {
		port = os.Getenv("MYSQLPORT")
	}
	if databaseName == "" {
		databaseName = os.Getenv("MYSQLDATABASE")
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

	if err := db.AutoMigrate(&note.Note{}); err != nil {
		return db, err
	}

	return db, nil
}
