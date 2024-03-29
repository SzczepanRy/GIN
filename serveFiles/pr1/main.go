package main

import (
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	router := gin.Default()
// 	router.LoadHTMLGlob("static/*")
// 	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", gin.H{
// 			"title": "Main website",
// 		})
// 	})
// 	router.GET("/index.js", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.js", gin.H{
// 			"title": "Mainjs",
// 		})
// 	})
// 	router.Run(":3000")
// }

func serveFileHeader(url string) string {
	fmt.Println(url)
	switch path.Ext(url) {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	default:
		return "text/plain"
	}

}

func main() {
	r := gin.Default()
	// r.LoadHTMLGlob("static/*")
	r.Static("/", "./static")
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	// r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// r.GET("/", func(c *gin.Context) {
	// 	fmt.Println(c.Request.URL.Path)
	// 	url := c.Request.URL.Path
	// 	filename := url[1:]
	// 	if filename != "" {
	// 		c.Header("Content-Type", serveFileHeader(url))
	// 		c.File(filename)
	// 	} else {
	// 		c.HTML(http.StatusOK, "index.html", gin.H{
	// 			"title": "Main website",
	// 		})
	// 	}

	// })

	panic(r.Run(":3000"))
}
