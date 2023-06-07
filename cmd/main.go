package cmd

import (
	"fmt"
	"net/http"
)

const port = ":3001"

func Runner() {
	server := http.NewServeMux()
	routes(server)
	fs := http.FileServer(http.Dir("templates/assets"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("(http://localhost:8080", port)
	err := http.ListenAndServe(port, server)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
}

func routes(server *http.ServeMux) {
	server.HandleFunc("/", indexHandlers)
	server.HandleFunc("/register", registerHandlers)
	server.HandleFunc("/admin", adminHandlers)
	server.HandleFunc("/api/login", loginPost)
	server.HandleFunc("/api/take-ban", adminPanel)
	server.HandleFunc("/api/register", CreateUser)
	server.HandleFunc("/api/adminPanel", adminPanel)
	server.HandleFunc("/api/catch-info-admin", sendInfoAdmin)
}
