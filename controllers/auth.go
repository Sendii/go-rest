package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetToken(c *gin.Context){
	var (
		user structs.User
		result gin.H
	)

	username := c.PostForm("username")
	password := c.PostForm("password")

	err := idb.DB.Where("Username = ? AND Password = ?", username, password).First(&user).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"message": "User tidak ditemukan",
		}
	}else{
		result = gin.H{
			"result": user,
			"api_code": "11111",
		}
	}
	c.JSON(http.StatusOK, result)
}