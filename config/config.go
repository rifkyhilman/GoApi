package config

import (
	"Goapi/structs"

	"github.com/jinzhu/gorm"
)

//Membuat koneksi ke database Mysql
func InDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
