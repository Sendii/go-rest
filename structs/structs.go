package structs

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}

type User struct {
	gorm.Model
	Username string
	Password string
	Jenkel string
	Numphone string
}