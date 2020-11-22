package main

import (
	"implementasi_MVC/app/middleware"
	"implementasi_MVC/app/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default()
	
	// add CORS
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"GET", "POST"}
	cfg.AllowHeaders = []string{"Authorization", "Origin", "Accept", "X-Requested-With", " Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	router.Use(cors.New(cfg))

	router.POST("/api/v1/account/add", controller.CreateAccount)
	router.POST("/api/v1/login",controller.Login)
	router.GET("/api/v1/account",middleware.Auth,controller.GetAccount)
	router.POST("/api/v1/transfer",middleware.Auth,controller.Transfer)
	router.POST("/api/v1/withdraw",middleware.Auth,controller.Withdraw)
	router.POST("/api/v1/deposit",middleware.Auth,controller.Deposit)
	router.POST("/api/v1/interest", middleware.Auth, controller.Interest)
	router.GET("/api/v1/mutasi", middleware.Auth, controller.GetAccountMutasi)

	router.Run(":8080")
}
