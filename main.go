package main

import (
	"implementasi_MVC/app/middleware"
	"implementasi_MVC/app/controller"
	// "implementasi_MVC/app/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	
	// db := model.DBinit()
	// inDB := &controller.InDB{DB: db}

	router := gin.Default()
	router.Use(cors.Default())

	// router.GET("/",inDB.CreateDB)

	router.POST("/api/v1/account/add", controller.CreateAccount)
	router.POST("/api/v1/login",controller.Login)
	router.GET("/api/v1/account",middleware.Auth,controller.GetAccount)
	router.POST("/api/v1/transfer",middleware.Auth,controller.Transfer)
	router.POST("/api/v1/withdraw",middleware.Auth,controller.Withdraw)
	router.POST("/api/v1/deposit",middleware.Auth,controller.Deposit)

	router.Run(":8080")
}