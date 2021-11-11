package main

import (
	"Ders16_MYSQL/1.Controller"
	"Ders16_MYSQL/2.auth"
	"Ders16_MYSQL/3.lib"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//author: gurbanli


func main(){
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session_id", store))

	authController := controller.AuthController{}
	productController := auth.ProductController{}
	adminController := lib.AdminController{}

	router.POST("/login", authController.Login)
	router.GET("/products", productController.GetProducts)
	router.GET("/getmyname", productController.GetMyUsername)
	router.GET("/admin", adminController.GetAdminPage)
	router.Run(":2222")
}