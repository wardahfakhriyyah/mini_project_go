package database

import (
	"fmt"
	"miniproject_go_wardahfdn/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func NewDatabase() (*Database, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/miniproject_wardahfdn?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}

func (db *Database) AutoMigrate() error {
	err := db.Conn.AutoMigrate(
		&model.User{},
		&model.Restaurant{},
		&model.Menu{},
		&model.Order{},
		&model.Payment{},
	)
	if err != nil {
		return err
	}

	fmt.Println("Auto Migration has been processed")

	return nil
}
