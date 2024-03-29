package main

import (
	"github.com/gin-gonic/gin"
	"jwt.com/api/controllers"
	"jwt.com/api/initializers"
	"jwt.com/api/middleware"
)

func init() {
	initializers.LoadEnvVerables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.ReqAuthc, controllers.Validate)
	r.Run(":3000")
}
