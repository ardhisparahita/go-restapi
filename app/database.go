package app

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_restapi")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
