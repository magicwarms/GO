package belajar_golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	DB, err := sql.Open("mysql","root:@tcp(localhost:3306)/db_faskes?parseTime=true")
	if(err != nil){
		panic(err)
	}
	// connection pooling DB
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)

	return DB;
}
