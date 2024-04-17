package controllers

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Sqldb struct {
	DB *sql.DB
}

func Upload(ctx *gin.Context) {

	fmt.Println("Successfully connected!")

	fmt.Println("sent")

	b64, tableName := getFileb64(ctx)

	if b64 == "" {
		errors.New("faled to parse or did not receve file")
	}
	//SAVE WHATS WRITTEN TO THE CONOLE TO A FILE
	// os.Stdout.Write(fileBytes)

	// err = os.WriteFile("text.jpg", fileBytes, 0644)
	// fmt.Println(imgBase64Str)
	// _, err = tempFile.Write(fileBytes)

	// if err != nil {
	// 	ctx.Error(err)
	// 	return
	// }
	// bytesString = bytesString[1 : len(bytesString)-1]

	db := ConnectToDB()
	defer db.Close()

	sqldb := Sqldb{db}
	records, err := sqldb.GetRecords(tableName)
	if err != nil {
		ctx.Error(err)
	}

	for _, record := range records {
		fmt.Println(record.id)
	}

	getCount, err := sqldb.GetRecordsCount(tableName)
	if err != nil {
		ctx.Error(err)
	}
	lastId := getCount.count

	err = sqldb.WriteToDB(tableName, lastId+1, b64)

	if err != nil {
		errors.New("faled to write to db")
	}

	fmt.Println("redir")

	defer ctx.Redirect(http.StatusMovedPermanently, "http://localhost:3000/")
	// ctx.String(http.StatusOK, fmt.Sprintf("files uploaded!"))

}

func getFileb64(ctx *gin.Context) (string, string) {
	fileHeader, err := ctx.FormFile("file")
	val := ctx.Request.Form["table"]
	tableName := val[0]
	// fmt.Println(val)
	if err != nil {
		ctx.Error(err)
		return "", ""
	}

	//Open received file
	csvFileToImport, err := fileHeader.Open()
	if err != nil {
		ctx.Error(err)
		return "", ""
	}
	defer csvFileToImport.Close()

	//Create temp file
	tempFile, err := os.CreateTemp("", fileHeader.Filename)
	if err != nil {
		ctx.Error(err)
		return "", ""
	}
	defer tempFile.Close()

	// fmt.Println(tempFile.Name())
	//Delete temp file after importing
	defer os.Remove(tempFile.Name())

	// //Write data from received file to temp file
	fileBytes, err := io.ReadAll(csvFileToImport)
	if err != nil {
		ctx.Error(err)
		return "", ""
	}

	b64 := base64.StdEncoding.EncodeToString(fileBytes)
	return b64, tableName
}

func ConnectToDB() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "a"
		dbname   = "postgres"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func (sqldb Sqldb) WriteToDB(tableName string, id int, b64 string) error {

	insertDynStmt := fmt.Sprintf(`insert into "%v" ("id", "file") values($1, $2)`, tableName)
	_, err := sqldb.DB.Exec(insertDynStmt, id, b64)
	if err != nil {
		fmt.Println(tableName, id)
		fmt.Println(err)
		return errors.New("could not wirte to db ")
	}
	return nil

}

type rowType struct {
	id   int
	file string
}

func (sqldb Sqldb) GetRecords(tableName string) ([]rowType, error) {

	//would it  be bett
	query := fmt.Sprintf(`select * from "%v"`, tableName)
	rows, err := sqldb.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("could not query")
	}
	defer rows.Close()
	values := make([]rowType, 0)

	for rows.Next() {
		value := rowType{}
		err := rows.Scan(&value.id, &value.file)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("could not read row")
		}
		values = append(values, value)
	}
	// fmt.Println(values)
	fmt.Println("number of rows ", len(values))
	// num, err := strconv.Atoi(r)
	return values, nil
}

type rowCountType struct {
	count int
}

func (sqldb Sqldb) GetRecordsCount(tableName string) (rowCountType, error) {

	//would it  be bett
	query := fmt.Sprintf(`select count(*) from "%v"`, tableName)
	rows, err := sqldb.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return rowCountType{}, errors.New("could not query")
	}
	defer rows.Close()
	values := make([]rowCountType, 0)

	for rows.Next() {
		value := rowCountType{}
		err = rows.Scan(&value.count)
		if err != nil {
			fmt.Println(err)
			return rowCountType{}, errors.New("could not read row")
		}
		values = append(values, value)
	}
	fmt.Println("number of rows ", values[0])
	// num, err := strconv.Atoi(r)
	return values[0], nil
}
