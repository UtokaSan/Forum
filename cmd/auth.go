package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	_ "github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"

	"io/ioutil"
	"net/http"
	"time"
)

func loginPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var userLogin Login
	err = json.Unmarshal(body, &userLogin)
	token := jwt.New(jwt.SigningMethodHS256)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userFound := false
	for _, user := range readUsers() {
		if user.Email == userLogin.Email && userLogin.Password == user.Password {
			claim := token.Claims.(jwt.MapClaims)
			claim["user-id"] = user.ID
			claim["exp"] = time.Now().Add(time.Hour * 24).Unix()
			tokenStr, err := token.SignedString([]byte("token-user"))
			if err != nil {
				fmt.Println(err)
				return
			}
			cookieOrSession(w, r, userLogin.SaveInfo, tokenStr)
			w.WriteHeader(http.StatusOK)
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func cookieOrSession(w http.ResponseWriter, r *http.Request, userlogin string, tokenStr string) {
	if userlogin == "on" {
		cookie := http.Cookie{
			Name:    "jwtToken",
			Value:   tokenStr,
			Expires: time.Now().Add(time.Hour * 24),
			Path:    "/",
		}
		http.SetCookie(w, &cookie)
	} else {
		var store = sessions.NewCookieStore([]byte("secret-key"))
		session, err := store.Get(r, "session-login")
		if err != nil {
			fmt.Println(err)
		}
		session.Values["jwtToken"] = tokenStr
		err = session.Save(r, w)
	}
}
