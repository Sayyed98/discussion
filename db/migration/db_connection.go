package migration

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // for MySQL driver
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", ("root:password@tcp(127.0.0.1:3306)/discussions"))
	fmt.Println("error ", err)
	if err != nil {
		log.Fatal("connection error")
	}
}

func Ping() error {
	return db.Ping()
}
