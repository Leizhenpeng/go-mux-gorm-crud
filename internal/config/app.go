package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() *gorm.DB {
	d, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_URL")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db = d
	return nil
}

func GetDB() *gorm.DB {
	return db
}
