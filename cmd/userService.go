package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func createUser(user User) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	createUser := `INSERT INTO users (id, nom, email, password) VALUES (?, ?, ?, ?)`
	_, errCreate := db.Exec(createUser, user.ID, user.Username, user.Email, user.Password)
	if errCreate != nil {
		fmt.Println(err)
	}
	fmt.Println("User created successfully")
}
func readUser() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT id, nom, email, password FROM users"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Id : " + strconv.Itoa(user.ID) + " Username : " + user.Username + " Email : " + user.Email + " Password : " + user.Password)
	}
}

func updateUser(user User) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "UPDATE users SET nom = ?, email = ?, password = ? WHERE ID = ?"
	_, err = db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User update successfully")
}
func deleteUser(idOfUser int) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "DELETE FROM users WHERE id = ?"
	_, err = db.Exec(query, idOfUser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User delete successfully")
}
