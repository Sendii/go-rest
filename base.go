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

	p := r.Group("person")
	{
		p.GET("/", inDB.GetPersons)	
		p.POST("/create", inDB.CreatePerson) //use validation firstname and lastname cannot null
		p.GET("/:id", inDB.GetPerson)
		p.PUT("/update", inDB.UpdatePerson)

		//delete bisa dengan .DELETE ataupun .POST
		p.POST("/delete/:id", inDB.DeletePerson)
	}

	r.Run()
}