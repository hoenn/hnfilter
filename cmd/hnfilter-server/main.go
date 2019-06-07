package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/hoenn/hnfilter/pkg/store"
	_ "github.com/lib/pq"
)

func main() {
	//TODO move this to config, start http server
	dbConn := &store.DBConn{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
	db, err := sql.Open("postgres", dbConn.Format())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	ds := store.NewDataStore(db)
	fmt.Println(ds)
}
