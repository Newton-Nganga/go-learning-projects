package models


import (
	"gorm.io/gorm"
	_ "gorm.io/driver/sqlite"
)


var DB *gorm.DB



func ConnectDatabase(){
	database,err := gorm.Open(sqlite.Open("test.db"), &gorm.Config)



	if err != nil {
		panic("failed to connect database")
	}

	err = database.AutoMigrate(&Book{})

	if err != nil {
		return
	}

	DB = database
}