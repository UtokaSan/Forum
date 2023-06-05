package cmd

import (
	"fmt"
	"net/http"
)

const port = ":3000"

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", indexHandlers)
	//server.HandleFunc("/register", registerHandlers)
	server.HandleFunc("/api/login", loginPost)
	server.HandleFunc("/api/register", CreateUser)
	fs := http.FileServer(http.Dir("templates/assets"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("\n\033[34m[http://127.0.0.1:3000]\033[32m \033[4mServer run on port", port[1:], ".\033[0m")

	err := http.ListenAndServe("127.0.0.1:3000", server)

	if err != nil {
		fmt.Println("error :", err)
		return
	}
}
