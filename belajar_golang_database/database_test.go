package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	
}

func TestOpenConnection(t *testing.T) {
	DB, err := sql.Open("mysql","root:@tcp(localhosst:3306)/db_faskes")
	if(err != nil){
		panic(err)
	}
	defer DB.Close()
}