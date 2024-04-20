package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Id int `json:"id"`
}

func Download(ctx *gin.Context, db *sql.DB) {

	filesChan := make(chan string)
	finishChan := make(chan int, 10)

	var reqBody Body
	if err := ctx.BindJSON(&reqBody); err != nil {
		errors.New("could not parse req")
	}

	fmt.Println(reqBody)

	// db := ConnectToDB()
	// defer db.Close()
	sqldb := Sqldb{db}

	go sqldb.GetRecordsById("table1", reqBody.Id, filesChan, finishChan)
	go sqldb.GetRecordsById("table2", reqBody.Id, filesChan, finishChan)
	go sqldb.GetRecordsById("table3", reqBody.Id, filesChan, finishChan)

	// var id int
	// var fileB64 string

	var files []string
	fmt.Println("for started")

	var temp int = 0
	for {
		select {
		case file := <-filesChan:
			files = append(files, file)
		case <-finishChan:
			temp += 1
			if temp == 3 {
				ctx.JSON(http.StatusOK, gin.H{
					"id":   0,
					"file": files,
				})
				return
			}

		}

	}

}

func (sqldb Sqldb) GetRecordsById(tableName string, id int, filesChan chan<- string, finishChan chan<- int) {

	//would it  be bett
	query := fmt.Sprintf(`select * from "%v" where id=%v`, tableName, id)
	fmt.Println(query)
	rows, err := sqldb.DB.Query(query)

	if err != nil {
		fmt.Println("finish faled to query")
		finishChan <- 1
		fmt.Println(err)
		//return nil, errors.New("could not query")
	}
	defer rows.Close()
	//	values := make([]rowType, 0)

	for rows.Next() {
		value := rowType{}
		err := rows.Scan(&value.id, &value.file)
		if err != nil {
			fmt.Println("finish faled to scan")
			fmt.Println(err)
			finishChan <- 1
			//	return nil, errors.New("could not read row")
		}
		filesChan <- value.file
		// values = append(values, value)
	}
	fmt.Println("finish")
	finishChan <- 1

	// fmt.Println(values)
	// fmt.Println("number of rows ", len(values))
	// num, err := strconv.Atoi(r)
	// return values, nil
}
