package controllers

import (
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "path/filepath"
	"os"
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

func UploadFile(c *gin.Context){
	var result gin.H
	// Input Tipe Text
	nama := c.PostForm("nama")

	// Multiple Form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}
	// Files
	files := form.File["files"]
	fmt.Println(form)
	// For range
	imagepath := "./upload/images/"
	os.MkdirAll(imagepath, 0777)
	for _, file := range files {
		path := imagepath + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}
	}

	// Response
	result = gin.H{
		"message": "Berhasil mengupload file",
		"total": len(files),
		"nama": nama,
	}
	// c.String(http.StatusOK, fmt.Sprintf("Files count : %d, nama : %s", len(files), nama))
	c.JSON(http.StatusOK, result)
}