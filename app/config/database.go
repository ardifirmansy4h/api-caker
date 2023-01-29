package config

import (
	"caker/helper"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	var dtbs string = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helper.GetConfig("DB_USERNAME"),
		helper.GetConfig("DB_PASSWORD"),
		helper.GetConfig("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dtbs), &gorm.Config{})

	if err != nil {
		fmt.Println("Database tidak terkoneksi")
	}
	MigrateDB()
	fmt.Println("Database terkoneksi")
}
