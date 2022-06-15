package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	DB := GetConnection()
	defer DB.Close()

	ctx := context.Background()

	script := "INSERT INTO customers(id, name) VALUE('rani','Rani')"
	// untuk insert apa ExecContext
	_, err := DB.ExecContext(ctx, script)
	if(err != nil){
		panic(err)
	}

	fmt.Println("Insert customer data success")
}

func TestExecSelectSql(t *testing.T) {
	DB := GetConnection()
	// wajib close setelah tak di pakai
	defer DB.Close()
	ctx := context.Background()

	script := "SELECT * from customers"
	rows, err := DB.QueryContext(ctx, script)
	if(err != nil){
		panic(err)
	}
	// setelah rows di dapat atau gak error mesti kudu di tutup/close
	defer rows.Close()
	// untuk menampilkan data dari query diatas
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if(err != nil){
			panic(err)
		}
		println("ini id:", id)
		println("ini name:",name)
	}

}

func TestExecComplexSql(t *testing.T) {
	DB := GetConnection()
	// wajib close setelah tak di pakai
	defer DB.Close()
	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at from customers"
	rows, err := DB.QueryContext(ctx, script)
	if(err != nil){
		panic(err)
	}

	// untuk menampilkan data dari query diatas loh
	for rows.Next() {
		var id, name string
		// pakai sql.NullString kalau salah 1 column nya bisa null
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if(err != nil){
			panic(err)
		}
		fmt.Println("=======================================")
		fmt.Println("ini id: ", id)
		fmt.Println("ini name: ",name)
		// beri kondisi kalo dia valid tampilkan data nya otherwise tampilkan nil aja
		if(email.Valid) {
			fmt.Println("ini email: ", email.String)
		} else {
			fmt.Println("ini email: ",nil)
		}

		fmt.Println("ini balance: ", balance)
		fmt.Println("ini rating: ", rating)
		fmt.Println("ini birthDate: ", birthDate)
		fmt.Println("ini married: ", married)
		fmt.Println("ini createdAt: ", createdAt)
	}
	// setelah rows di dapat atau gak error mesti kudu di tutup/close
	defer rows.Close()
}
