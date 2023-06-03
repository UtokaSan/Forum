package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func createPost(post Post) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	postCreate := `INSERT INTO posts (id, photo, texte, hidden, like, dislike, report, categorie, ban) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, errCreate := db.Exec(postCreate, post.ID, post.Photo, post.Texte, post.Hidden, post.Like, post.Dislike, post.Signalement, post.Categorie, post.Ban)
	if errCreate != nil {
		fmt.Println(err)
	}
	fmt.Println("Post created successfully")
}
func readPost() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT id, photo, texte, hidden, like, dislike, report, categorie, ban FROM posts"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Photo, &post.Texte, &post.Hidden, &post.Like, &post.Dislike, &post.Signalement, &post.Categorie, &post.Ban)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Id : " + strconv.Itoa(post.ID) + " photo : " + post.Photo + " texte : " + post.Texte)
	}
}

func updatePost(post Post) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "UPDATE posts SET photo = ?, texte = ?, hidden = ?, like = ?, dislike = ?, report = ?, categorie = ?, ban = ? WHERE ID = ?"
	_, err = db.Exec(query, post.Photo, post.Texte, post.Hidden, post.Like, post.Dislike, post.Signalement, post.Categorie, post.Ban)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User update successfully")
}
func deletePost(idOfUser int) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "DELETE FROM posts WHERE id = ?"
	_, err = db.Exec(query, idOfUser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User delete successfully")
}
