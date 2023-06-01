package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func loginPost(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT * FROM users WHERE nom = ?"
	rows, err := db.Query(query, "snake")
	if err != nil {
		fmt.Println(err)
	}
	if rows.Next() {
		fmt.Println("User exist")
	} else {
		fmt.Println("don't exist")
	}
}
