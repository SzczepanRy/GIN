package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clinets []websocket.Conn

func main() {

	// r := gin.Default()

	// r.GET("/",)

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		clinets = append(clinets, *conn)

		//main loop
		for {
			//read
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))
			//loop if found and sent to browser
			for _, client := range clinets {
				if err = client.WriteMessage(msgType, msg); err != nil {
					return
				}

			}
		}

	})

	//send files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
		//write and delete index
	})
	println("running on 3000")

	http.ListenAndServe(":3000", nil)

	// r.Static("/", "./static")
	// panic(r.Run(":3000"))
}
