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
	const secretToken = "token-user"
	token := getSession(r)
	tokenJWT := checkJWT(secretToken, token)
	dataUser := getData(tokenJWT)

	fmt.Println(dataUser)
	jsonData, err := json.Marshal(unhiddenPost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataUser)
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

func takePostWithId(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := TakePostId{
		Title: "ss",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
	w.Write(jsonData)
}

func postLikeOrDislike(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var data Reaction
	err = json.Unmarshal(body, &data)
	if data.Like == "on" {
		if likePost(data.UserId, data.PostId) == true {
			w.Write([]byte("liked"))
		} else {
			w.Write([]byte("already liked"))
		}
	}
	if data.Dislike == "on" {
		if dislikePost(data.UserId, data.PostId) == true {
			w.Write([]byte("liked"))
		} else {
			w.Write([]byte("already liked"))
		}
	}
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
