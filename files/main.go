package main

import (
	"time"

	"files.com/files/conn"
	"files.com/files/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// type Sqldb struct {
// 	DB *sql.DB
// }

func main() {

	db := conn.ConnectToDB()
	defer db.Close()
	// sqldb := Sqldb{db}
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Static("/", "./static")
	r.POST("/upload", func(ctx *gin.Context) {
		controllers.Upload(ctx, db)
	})
	r.POST("/get", func(ctx *gin.Context) {
		controllers.Download(ctx, db)
	})

	r.Run(":3000")

}
