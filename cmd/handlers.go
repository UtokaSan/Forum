package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"html/template"
	"io/ioutil"
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
	callbackLoginGoogle(w, r)

	//if r.URL.Path != "/admin" {
	//	errorHandler(w, r, http.StatusNotFound)
	//} else {
	//	t, err := template.ParseFiles("templates/Admin.html")
	//	takeInfoGoogle(r)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	t.Execute(w, r)
	//}
}

func takeInfoGoogle(r *http.Request) {
	config := getConfig()
	code := r.URL.Query().Get("code")
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := config.Client(oauth2.NoContext, token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(response.Body)

	var usergoogle UserGoogle
	err = json.Unmarshal(body, &usergoogle)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(usergoogle)

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
