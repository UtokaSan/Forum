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
	createUser := `INSERT INTO users (id, image, nom, email, password) VALUES (?, ?, ?, ?, ?)`
	_, errCreate := db.Exec(createUser, user.ID, user.Image, user.Username, user.Email, user.Password)
	if errCreate != nil {
		fmt.Println(err)
	}
	fmt.Println("User created successfully")
}
func readUser() []User {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT id, image, nom, email, password FROM users"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	var tab []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Image, &user.Username, &user.Email, &user.Password)
		if err != nil {
			fmt.Println(err)
		}
		result := append(tab, user)
		fmt.Println("Id : " + strconv.Itoa(user.ID) + " Username : " + user.Username + " Email : " + user.Email + " Password : " + user.Password)
		return result
	}
	return tab
}

func updateUser(user User) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "UPDATE users SET nom = ?, image = ?, email = ?, password = ? WHERE ID = ?"
	_, err = db.Exec(query, user.Username, user.Image, user.Email, user.Password)
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
