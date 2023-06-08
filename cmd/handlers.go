package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"html/template"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	takeCookie, err := r.Cookie("jwtToken")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	token, err := jwt.Parse(takeCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Méthode de signature inattendue : %v", token.Header["alg"])
		}
		return []byte("token-user"), nil
	})
	if err != nil || !token.Valid {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if token.Valid {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func loginHandlers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
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

func adminHandlers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("templates/Admin.html")
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
