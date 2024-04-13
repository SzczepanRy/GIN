package main

import (
	"log"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var updrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnect(ctx *gin.Context) {

	jobs := make(chan int, 1000)
	results := make(chan int, 1000)
	conn, err := updrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Print("could not upgrade connectioin")
		ctx.AbortWithError(http.StatusInternalServerError, nil)
		return
	}

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	for i := 0; i < 1000; i++ {
		jobs <- i
	}
	close(jobs)
	for {

		bs := []byte(strconv.Itoa(<-results))
		if err = conn.WriteMessage(websocket.TextMessage, bs); err != nil {
			log.Print("could not send to client")
			//ctx.AbortWithError(http.StatusInternalServerError, nil)
			return
		}

	}

}
func worker(jobs <-chan int, results chan int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n == 0 {
		return -1
	}
	if n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {

	r := gin.Default()
	r.Static("/static", "./static")

	r.GET("/connect", handleConnect)

	r.Run(":3000")

}
