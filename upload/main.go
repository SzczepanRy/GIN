package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type typeVal struct {
	id  string
	val string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/id/{id}", func(w http.ResponseWriter, r *http.Request) {
		userCh := make(chan typeVal)
		dbCh := make(chan typeVal)
		testCh := make(chan typeVal)

		endCh := make(chan []typeVal)
		id := r.PathValue("id")
		log.Print(id)

		go getUser(userCh)
		go getDB(dbCh)
		go getTest(testCh)
		go getdata(id, userCh, dbCh, testCh, endCh)

		val := <-endCh
		w.Write([]byte(fmt.Sprintf("%#v", val)))

	})

	log.Print("running on 3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getdata(id string, ch1 chan typeVal, ch2 chan typeVal, ch3 chan typeVal, endCh chan []typeVal) []typeVal {
	var endData []typeVal
	for {
		wart1 := <-ch1
		wart2 := <-ch2
		wart3 := <-ch3
		if wart1.id == id {

			endData = append(endData, wart1)
		}
		if wart2.id == id {

			endData = append(endData, wart2)
		}
		if wart2.id == id {

			endData = append(endData, wart3)
		}
		// log.Print(wart.id, " : ", wart.val)
		if len(endData) == 3 {
			endCh <- endData

		}
	}

}

func getUser(ch chan typeVal) {
	dat, err := os.ReadFile("./users.txt")
	if err != nil {
		log.Print("BAD READ")
	}
	data := string(dat)

	datasl := strings.Split(data, "\n")

	for _, val := range datasl {
		innerdata := strings.Split(val, " ")
		ch <- typeVal{id: innerdata[0], val: innerdata[len(innerdata)-1]}
		// endData[innerdata[0]] = innerdata[1]

	}

}
func getDB(ch chan typeVal) {
	dat, err := os.ReadFile("./users.txt")
	if err != nil {
		log.Print("BAD READ")
	}
	data := string(dat)

	datasl := strings.Split(data, "\n")

	for _, val := range datasl {
		innerdata := strings.Split(val, " ")
		ch <- typeVal{id: innerdata[0], val: innerdata[len(innerdata)-1]}
		// endData[innerdata[0]] = innerdata[1]

	}

}
func getTest(ch chan typeVal) {
	dat, err := os.ReadFile("./users.txt")
	if err != nil {
		log.Print("BAD READ")
	}
	data := string(dat)

	datasl := strings.Split(data, "\n")

	for _, val := range datasl {
		innerdata := strings.Split(val, " ")
		ch <- typeVal{id: innerdata[0], val: innerdata[len(innerdata)-1]}
		// endData[innerdata[0]] = innerdata[1]

	}

}
