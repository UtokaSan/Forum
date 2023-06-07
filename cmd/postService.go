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
	postCreate := `INSERT INTO posts (id, photo, texte, hidden, like, dislike, report, categorie, ban, archived) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, errCreate := db.Exec(postCreate, post.ID, post.Photo, post.Texte, post.Hidden, post.Like, post.Dislike, post.Signalement, post.Categorie, post.Ban, post.Archived)
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
	query := "SELECT id, photo, texte, hidden, like, dislike, report, categorie, ban, archived FROM posts"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Photo, &post.Texte, &post.Hidden, &post.Like, &post.Dislike, &post.Signalement, &post.Categorie, &post.Ban, &post.Archived)
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
	query := "UPDATE posts SET photo = ?, texte = ?, hidden = ?, like = ?, dislike = ?, report = ?, categorie = ?, ban = ?, archived = ? WHERE ID = ?"
	_, err = db.Exec(query, post.Photo, post.Texte, post.Hidden, post.Like, post.Dislike, post.Signalement, post.Categorie, post.Ban, post.Archived)
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

func takePostHidden() []map[string]interface{} {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT * FROM posts WHERE hidden = 1"
	rows, err := db.Query(query)
	var result []map[string]interface{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Photo, &post.Texte, &post.Hidden, &post.Like, &post.Dislike, &post.Signalement, &post.Categorie, &post.Ban, &post.Archived)
		if err != nil {
			fmt.Println(err)
		}
		postData := make(map[string]interface{})
		postData["id"] = post.ID
		postData["texte"] = post.Texte
		postData["categorie"] = post.Categorie
		result = append(result, postData)
	}
	return result
}

func postArchived() []map[string]interface{} {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT * FROM posts WHERE archived = 1"
	rows, err := db.Query(query)
	var result []map[string]interface{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Photo, &post.Texte, &post.Hidden, &post.Like, &post.Dislike, &post.Signalement, &post.Categorie, &post.Ban, &post.Archived)
		if err != nil {
			fmt.Println(err)
		}
		postData := make(map[string]interface{})
		postData["id"] = post.ID
		postData["texte"] = post.Texte
		postData["categorie"] = post.Categorie
		result = append(result, postData)
	}
	return result
}
