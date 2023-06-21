package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	const secretToken = "token-user"
	token := getSession(r)
	tokenJWT := checkJWT(secretToken, token)
	dataUser := getData(tokenJWT)
	if err != nil {
		fmt.Println(err)
		return
	}
	var data GestionPost
	err = json.Unmarshal(body, &data)
	createPostWithTitle(Post{
		Title:     data.CreatePost,
		IDCreator: dataUser.UserId,
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
			w.Write([]byte("already disliked"))
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
			println("C'est de la merde")
			return
		}
		updatePost(postEdit)
	} else if dataUser.UserId == post.IDCreator {
		fmt.Println("user")
		postEdit := changedDataPost(post, data)
		if postEdit.ID == -1 {
			println("C'est de la merde")
			return
		}
		updatePost(postEdit)

	} else {

		fmt.Println("Mec user Other")

		// REFUSE
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
