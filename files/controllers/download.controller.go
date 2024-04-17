package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Id int `json:"id"`
}

func Download(ctx *gin.Context) {

	filesChan := make(chan string)
	finishChan := make(chan int, 3)

	var reqBody Body
	if err := ctx.BindJSON(&reqBody); err != nil {
		errors.New("could not parse req")
	}

	fmt.Println(reqBody)

	db := ConnectToDB()
	defer db.Close()
	sqldb := Sqldb{db}

	go sqldb.GetRecordsById("table1", reqBody.Id, filesChan, finishChan)
	go sqldb.GetRecordsById("table2", reqBody.Id, filesChan, finishChan)
	go sqldb.GetRecordsById("table3", reqBody.Id, filesChan, finishChan)

	// var id int
	// var fileB64 string

	var files []string

	for {

		file := <-filesChan
		files = append(files, file)
		fmt.Println(len(finishChan))

		if len(finishChan) == 2 {
			fmt.Println("sent")
			ctx.JSON(http.StatusOK, gin.H{
				"id":   0,
				"file": files,
			})
			return
		}
	}

}

// for _, file := range records {
// 	id = file.id
// 	fileB64 = file.file
// }

func (sqldb Sqldb) GetRecordsById(tableName string, id int, filesChan chan<- string, finishChan chan<- int) {

	//would it  be bett
	query := fmt.Sprintf(`select * from "%v" where id=%v`, tableName, id)
	fmt.Println(query)
	rows, err := sqldb.DB.Query(query)
	if err != nil {
		fmt.Println("finish")
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
			fmt.Println("finish")
			finishChan <- 1
			fmt.Println(err)
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
