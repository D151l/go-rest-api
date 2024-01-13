package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var databaseConnection *sql.DB

func Connect() {
	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	databaseConnection = db
}

func CreateTables() {
	var tableQuery = "create table todos (id varchar(30) not null, item varchar(32) not null, completed bit not null, primary key (id))"
	_, err := databaseConnection.Exec(tableQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Close() {
	if databaseConnection != nil {
		err := databaseConnection.Close()
		if err != nil {
			return
		}
	}
}

func AddTodo(todo todo) {

}
