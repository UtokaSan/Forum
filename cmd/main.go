package cmd

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", loginHandlers)
	server.HandleFunc("/register", registerHandlers)
	server.HandleFunc("/middleware", authAdminSecurity(adminHandlers))
	server.HandleFunc("/api/login", loginPost)
	server.HandleFunc("/api/register", CreateUser)
	fs := http.FileServer(http.Dir("templates/assets"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("\n\033[34m[http://127.0.0.1:8080]\033[32m \033[4mServer run on port", port[1:], ".\033[0m")

	err := http.ListenAndServe("127.0.0.1:8080", server)

	if err != nil {
		fmt.Println("error :", err)
		return
	}
}
