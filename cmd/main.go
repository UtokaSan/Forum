package cmd

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Runner() {
	server := http.NewServeMux()
	routes(server)
	fs := http.FileServer(http.Dir("templates/assets"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("\n\033[34m[http://127.0.0.1"+port+"]\033[32m \033[4mServer run on port", port, ".\033[0m")

	err := http.ListenAndServe(port, server)

	if err != nil {
		fmt.Println("error :", err)
		return
	}
}

func routes(server *http.ServeMux) {
	//server.HandleFunc("/", authGuestSecurity(rootHandler))
	server.HandleFunc("/homepage", authGuestSecurity(mainHandlers))
	server.HandleFunc("/login", authGuestSecurity(loginHandlers))
	server.HandleFunc("/login/google", authGuestSecurity(loginGoogle))
	server.HandleFunc("/login/github", authGuestSecurity(loginGithub))
	server.HandleFunc("/register", authGuestSecurity(registerHandlers))
	server.HandleFunc("/admin", authAdminSecurity(adminHandlers))
	server.HandleFunc("/post", authGuestSecurity(postHandlers))
	server.HandleFunc("/api/login", loginPost)
	//server.HandleFunc("/api/test", CreateAccountGoogle)
	server.HandleFunc("/api/loginGoogle", loginGoogle)
	server.HandleFunc("/api/loginGithub", loginGithub)
	server.HandleFunc("/api/callbacklogingoogle", callbackLoginGoogle)
	server.HandleFunc("/api/callbacklogingithub", callbackLoginGithub)
	server.HandleFunc("/api/register", CreateUser)
	server.HandleFunc("/api/adminpanel", authAdminSecurity(adminPanel))
	server.HandleFunc("/api/catch-info-admin", authAdminSecurity(sendInfoAdmin))
	server.HandleFunc("/api/create-post", authUserSecurity(createPostHandler))
	server.HandleFunc("/api/display-post", displayPostVisible)
	server.HandleFunc("/api/createcomment", authUserSecurity(createComment))
	server.HandleFunc("/api/editPost", authUserSecurity(editPost))
	server.HandleFunc("/api/getComments", authGuestSecurity(getComments))
	server.HandleFunc("/api/takepostid", authGuestSecurity(sendDataPostWithId))
	server.HandleFunc("/api/likeordislike", authUserSecurity(postLikeOrDislike))
}
