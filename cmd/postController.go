package cmd

import (
	"database/sql"
	"fmt"
)

func deletePost(post Post) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "DELETE FROM posts WHERE title = ?"
	_, err = db.Exec(query, post.Title)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Post delete successfully")
}

func createPostWithTitle(post Post) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	postCreate := `INSERT INTO posts (title) VALUES (?)`
	_, errCreate := db.Exec(postCreate, post.Title)
	if errCreate != nil {
		fmt.Println(err)
	}
	fmt.Println("Post created successfully")
}
