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
	postCreate := `INSERT INTO posts (photo, title, texte, hidden, like, dislike, report, categorie, ban, archived) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, errCreate := db.Exec(postCreate, post.Photo, post.Texte, post.Hidden, post.Like, post.Dislike, post.Signalement, post.Categorie, post.Ban, post.Archived)
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
	query := "SELECT id, photo, title, texte, hidden, like, dislike, report, categorie, ban, archived FROM posts"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Photo, &post.Title, &post.Texte, &post.Hidden, &post.Like, &post.Dislike, &post.Signalement, &post.Categorie, &post.Ban, &post.Archived)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Id : " + strconv.Itoa(post.ID) + " photo : " + post.Photo + " titre : " + post.Title + " texte : " + post.Texte)
	}
}

func readOnePostById() {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT id, photo, title, texte, hidden, like, dislike, signalement, categorie, ban, archived FROM users WHERE id = ?"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Photo, &post.Title, &post.Texte, &post.Hidden, &post.Like, &post.Dislike, &post.Signalement, &post.Categorie, &post.Ban, &post.Archived)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Id : " + strconv.Itoa(post.ID) + " photo : " + post.Photo + " titre : " + post.Title + " texte : " + post.Texte)
	}
}

func updatePost(post Post) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "UPDATE posts SET photo = ?, title = ?, texte = ?, hidden = ?, like = ?, dislike = ?, report = ?, categorie = ?, ban = ?, archived = ? WHERE ID = ?"
	_, err = db.Exec(query, post.Photo, post.Title, post.Texte, post.Hidden, post.Like, post.Dislike, post.Signalement, post.Categorie, post.Ban, post.Archived)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User update successfully")
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
		err := rows.Scan(&post.ID, &post.Photo, &post.Title, &post.Texte, &post.Hidden, &post.Like, &post.Dislike, &post.Signalement, &post.Categorie, &post.Ban, &post.Archived)
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

func takePostUnHidden() []map[string]interface{} {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "SELECT * FROM posts WHERE hidden = 0"
	rows, err := db.Query(query)
	var result []map[string]interface{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Photo, &post.Title, &post.Texte, &post.Hidden, &post.Like, &post.Dislike, &post.Signalement, &post.Categorie, &post.Ban, &post.Archived)
		if err != nil {
			fmt.Println(err)
		}
		postData := make(map[string]interface{})
		postData["id"] = post.ID
		postData["title"] = post.Title
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
		err := rows.Scan(&post.ID, &post.Photo, &post.Title, &post.Texte, &post.Hidden, &post.Like, &post.Dislike, &post.Signalement, &post.Categorie, &post.Ban, &post.Archived)
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

func createCommentService(comment Comment) {
	db, err := sql.Open("sqlite3", "./forum.db")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	postCreate := `INSERT INTO comments (IDPost, IDCreator,Text) VALUES (?,?,?)`
	_, errCreate := db.Exec(postCreate, comment.IDPost, comment.IDCreator, comment.Text)

	if errCreate != nil {
		fmt.Println(err)
	}
	fmt.Println("Post created successfully")
}
