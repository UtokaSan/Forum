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
	createUser := `INSERT INTO users (image, nom, email, password, role, ban) VALUES (?, ?, ?, ?, ?, ?)`
	_, errCreate := db.Exec(createUser, user.Image, user.Username, user.Email, user.Password)
	if errCreate != nil {
		fmt.Println(err)
	}
	fmt.Println("User created successfully")
}
func readUsers() []User {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT id, image, nom, email, password, role, ban FROM users"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	var tab []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Image, &user.Username, &user.Email, &user.Password, &user.Role, &user.Ban)
		if err != nil {
			fmt.Println(err)
		}
		tab = append(tab, user)
		fmt.Println("Id : " + strconv.Itoa(user.ID) + " Username : " + user.Username +
			" Email : " + user.Email + " Password : " + user.Password + " Role : " + user.Role + " " +
			"Ban : " + strconv.Itoa(user.Ban))
	}
	return tab
}

func updateUser(user User) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "UPDATE users SET nom = ?, image = ?, email = ?, password = ?, role = ?, ban = ?, report = ? WHERE ID = ?"
	_, err = db.Exec(query, user.Username, user.Image, user.Email, user.Password, user.Role, user.Ban, user.Report)
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

func readUser(id int) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT id, nom, email, password, role, ban FROM users WHERE id = ?"
	row := db.QueryRow(query, id)
	var user User
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.Ban, &user.Report)
	if err != nil {
		fmt.Println(err)
	}
}

func takeUsersBan() []map[string]interface{} {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT * FROM users WHERE ban = 0"
	rows, err := db.Query(query)
	var result []map[string]interface{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Image, &user.Username, &user.Email, &user.Password, &user.Role, &user.Ban, &user.Report)
		if err != nil {
			fmt.Println(err)
		}
		userData := make(map[string]interface{})
		userData["email"] = user.Email
		userData["username"] = user.Username
		result = append(result, userData)
	}
	return result
}

func takeAllUsers() []map[string]interface{} {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	var mapUser []map[string]interface{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Image, &user.Username, &user.Email, &user.Password, &user.Role, &user.Ban, &user.Report)
		if err != nil {
			fmt.Println(err)
		}
		userData := make(map[string]interface{})
		userData["email"] = user.Email
		userData["username"] = user.Username
		userData["role"] = user.Role
		mapUser = append(mapUser, userData)
	}
	return mapUser
}

func takeUserReported() []map[string]interface{} {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT * FROM users WHERE report > 0"
	rows, err := db.Query(query)
	var result []map[string]interface{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Image, &user.Username, &user.Email, &user.Password, &user.Role, &user.Ban, &user.Report)
		if err != nil {
			fmt.Println(err)
		}
		userData := make(map[string]interface{})
		userData["email"] = user.Email
		userData["username"] = user.Username
		result = append(result, userData)
	}
	return result
}
