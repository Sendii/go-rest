package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

//get one data with {id}
func (idb *InDB) GetUser(c *gin.Context) {
	var (
		user structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

//get all data in user
func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		users []structs.User
		result gin.H
	)

	idb.DB.Find(&users)
	if len(users) <= 0 {
		result = gin.H{
			"result": nil,
			"count": 0,
		}
	}else{
		result = gin.H{
			"result": users,
			"count": len(users),
		}
		CreateLog()
	}
	c.JSON(http.StatusOK, result)
}

//create new data to the database
func (idb *InDB) CreateUser(c *gin.Context){
	var (
		user structs.User
		result gin.H
	)
	username := c.PostForm("addusername")
	password := c.PostForm("addpassword")
	jenkel := c.PostForm("addjenkel")
	nomor := c.PostForm("addnomor")
	user.Username = username
	user.Password = password
	user.Jenkel = jenkel
	user.Numphone = nomor
	if username != "" && password != "" && jenkel != "" && nomor != ""{		
		idb.DB.Create(&user)

		result = gin.H{
			"result": user,
		}
	}else{
		result = gin.H{
			"message": "nama tidak boleh kosong",
		}
	}
	c.JSON(http.StatusOK, result)
}

//update data with {id} as query
func (idb *InDB) UpdateUser(c *gin.Context){
	id := c.Query("id")
	username := c.PostForm("updateusername")
	password := c.PostForm("updatepassword")
	jenkel := c.PostForm("updatejenkel")
	nomor := c.PostForm("updatenomor")
	var (
		user structs.User
		newUser structs.User
		result gin.H
	)

	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"message": "data not found",
		}
	}

	newUser.Username = username
	newUser.Password = password
	newUser.Jenkel = jenkel
	newUser.Numphone = nomor
	err = idb.DB.First(&user).Updates(newUser).Error
	if err != nil {
		result = gin.H{
			"message": "Update failed",
		}
	}else{
		result = gin.H{
			"data": newUser,
			"message": "Update Success",
		}
	}
	c.JSON(http.StatusOK, result)
}

//delete data with {id}
func (idb *InDB) DeleteUser(c *gin.Context){
	var (
		user structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"message": "data not found",
		}
	}
	err = idb.DB.Delete(&user).Error
	if err != nil {
		result = gin.H{
			"message": "delete failed",
		}
	}else{
		result = gin.H{
			"message": "Data delete Success",
		}
	}
	c.JSON(http.StatusOK, result)
}