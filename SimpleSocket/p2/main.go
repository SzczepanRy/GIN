package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")

	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		//check if the protocol can become a websocket one
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Panic("could not create a ws connection")
		}
		clients = append(clients, *conn)
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))
			//loop if found and sent to browser
			for _, client := range clients {
				if err = client.WriteMessage(msgType, msg); err != nil {
					return
				}

			}
		}

	})

	log.Print("running on 3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
