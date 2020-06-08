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
	people := structs.Person{First_Name: "Sendi", Last_Name: "Dian"}
	user := structs.User{Username: "Sendi", Password: "888888", Jenkel: "L", Numphone: "081220201131"}

	db.NewRecord(people) // => returns `true` as primary key is blank
	db.NewRecord(user)

	db.Create(&people)
	db.Create(&user)
	return db
}