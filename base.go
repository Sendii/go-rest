package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	// Static
	r.Static("/a", "./static")

	p := r.Group("person")
	{
		p.GET("/", inDB.GetPersons)	
		p.POST("/create", inDB.CreatePerson) //use validation firstname and lastname cannot null
		p.GET("/:id", inDB.GetPerson)
		p.POST("/update", inDB.UpdatePerson)

		//delete bisa dengan .DELETE ataupun .POST
		p.POST("/delete/:id", inDB.DeletePerson)
	}

	u := r.Group("user")
	{
		u.GET("/", inDB.GetUsers)	
		u.POST("/create", inDB.CreateUser) 
		u.GET("/:id", inDB.GetUser)
		u.POST("/update", inDB.UpdateUser)

		//delete bisa dengan .DELETE ataupun .POST
		u.POST("/delete/:id", inDB.DeleteUser)
	}

	a:= r.Group("auth")
	{
		a.POST("/login", inDB.GetToken)
	}

	r.POST("/upload", controllers.UploadFile)

	r.Run()
}