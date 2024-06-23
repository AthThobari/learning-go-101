package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID        uint
	Name      string
	Age       int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewDatabase() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/mengenal_sql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Something went wrong in connection database", err)
		return nil
	}

	fmt.Println("Database Connected")

	db.AutoMigrate(&Student{})
	return db
}
