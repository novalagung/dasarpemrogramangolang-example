package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlNullGeneric() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var name sql.Null[string]
	var grade sql.Null[int]

	err = db.QueryRow("select name, grade from tb_student where id = ?", "E001").
		Scan(&name, &grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if name.Valid {
		fmt.Println("name:", name.V)
	} else {
		fmt.Println("name: NULL")
	}

	if grade.Valid {
		fmt.Println("grade:", grade.V)
	} else {
		fmt.Println("grade: NULL")
	}
}

func main() {
	sqlNullGeneric()
}
