package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB * gorm.DB
)

func DBConnection() {
	dsn := "root:021916964cC@/auth-mysql-go?charset=utf8&parseTime=True&loc=Local";

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{});	
	if (err != nil) {
		panic(err);
	}; 

	DB = connection;
}