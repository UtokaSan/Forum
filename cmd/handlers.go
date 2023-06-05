package cmd

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandlers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("templates/Login.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, r)
	}
}
func registerHandlers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("templates/Inscription.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, r)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("templates/404" + ".html")
		if err != nil {
			fmt.Println(err)
		} else {
			t.Execute(w, r)
		}
	}
}
