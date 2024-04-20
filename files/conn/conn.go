package conn

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Sqldb struct {
	DB *sql.DB
}

func conn() *sql.DB {

	db := ConnectToDB()

	return db
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
