package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	uploadFunction := uploadImage(w, r)
	var image string
	if strings.Contains(uploadFunction, "assets") {
		image = uploadFunction
	}
	const secretToken = "token-user"
	token := getSession(r)
	tokenJWT := checkJWT(secretToken, token)
	dataUser := getData(tokenJWT)
	createPost(Post{
		Categorie: r.FormValue("action"),
		Title:     r.FormValue("message"),
		Texte:     r.FormValue("messageContent"),
		Photo:     image,
		IDCreator: dataUser.UserId,
	})
	w.WriteHeader(http.StatusOK)
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

func sendDataPostWithId(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var data Input
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	gestionPost := TakePostId{
		Info: takeInfoPostId(data.ID),
	}
	jsonData, err := json.Marshal(gestionPost)
	if err != nil {
		fmt.Println(err)
	}
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
	fmt.Println("ma reactions: ", data.PostId)
	const secretToken = "token-user"
	token := getSession(r)
	tokenJWT := checkJWT(secretToken, token)
	dataUser := getData(tokenJWT)
	if data.Reactions == "like" {
		if likePost(dataUser.UserId, data.PostId) == true {
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	if data.Reactions == "dislike" {
		if dislikePost(dataUser.UserId, data.PostId) == true {
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func editPost(w http.ResponseWriter, r *http.Request) {
	const secretToken = "token-user"
	token := getSession(r)
	tokenJWT := checkJWT(secretToken, token)
	dataUser := getData(tokenJWT)

	data := getDataEditPost(r)
	post := readOnePostById(data.ID)

	if dataUser.UserRole >= 3 {
		fmt.Println("Admin")
		postEdit := changedDataPost(post, data)
		if postEdit.ID == -1 {
			println("it's no change post with data post")
			return
		}
		updatePost(postEdit)
	} else if dataUser.UserId == post.IDCreator {
		fmt.Println("user")
		postEdit := changedDataPost(post, data)
		if postEdit.ID == -1 {
			println("it's no change post with data post")
			return
		}
		updatePost(postEdit)

	} else {
		return
	}
}

func getComments(w http.ResponseWriter, r *http.Request) {
	data := getDataComments(r)
	if data.ID == -1 {
		w.WriteHeader(200)
		w.Write([]byte("Error to get data"))
		return
	}
	comments := takeComments(data.ID)

	jsonData, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
	w.Write(jsonData)
}
