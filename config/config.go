package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
)

//DBInit Create connect to database
func DBInit() *gorm.DB{
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/lat_go?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic("failed to connect database")
	}

	db.AutoMigrate(structs.Person{}, structs.User{})

	//create dummy data
	user := structs.Person{First_Name: "Sendi", Last_Name: "Dian"}

	db.NewRecord(user) // => returns `true` as primary key is blank

	db.Create(&user)
	return db
}