package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var data GestionPost
	err = json.Unmarshal(body, &data)
	createPostWithTitle(Post{
		Title: data.CreatePost,
	})
}

func displayPostVisible(w http.ResponseWriter, r *http.Request) {
	unhiddenPost := takePostUnHidden()
	jsonData, err := json.Marshal(unhiddenPost)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(jsonData)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create comment")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var data Comment
	err = json.Unmarshal(body, &data)

	fmt.Println(data)

	if createCommentController(data) {
		return // error
	}
	createCommentService(data)

}

func editPost(w http.ResponseWriter, r *http.Request) {

	//const secretToken = "token-user"
	//token := getSession(r)
	//tokenJWT := checkJWT(secretToken, token)
	//dataUser := getData(tokenJWT)

	var dataUser DataTokenJWT
	dataUser.UserRole = 3
	dataUser.UserId = 2

	post := readOnePostById(2)

	if dataUser.UserRole >= 3 {
		fmt.Println("Admin")
		postEdit := editedPost(r, post)
		if postEdit.ID == -1 {
			println("C'est de la merde")
			return
		}
		fmt.Println("postEdit")
		updatePost(postEdit)

		fmt.Println(postEdit)

	} else if dataUser.UserId == 2 {
		fmt.Println("User")

	} else {

		fmt.Println("Mec user Other")

		// REFUSE
	}
}
