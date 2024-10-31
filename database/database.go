package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (db *gorm.DB, err error) {
	dsn := "host=localhost user=postgres password=chatApp dbname=core port=24031 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("database connected to : %v", dsn)
	return db, err
}
