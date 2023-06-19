package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"time"
)

func loginGoogle(w http.ResponseWriter, r *http.Request) {
	config := getConfig("116188844729-bpmpofo72u5vdhdt43qif41lmppqejuh.apps.googleusercontent.com", "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", []string{"https://www.googleapis.com/auth/userinfo.email"})
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func loginGithub(w http.ResponseWriter, r *http.Request) {
	config := getConfig("5ec9ece005f647affa3d", "df2a2f79c65ae75448eb034ae6d918d987e8973d", []string{"user:email"})
	url := config.AuthCodeURL("state", oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func getConfig(clientID string, clientSecret string, auth []string) *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/admin",
		Scopes:       auth,
		Endpoint:     google.Endpoint,
	}
	return config
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var userLogin Login
	err = json.Unmarshal(body, &userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users := readUsers()
	for _, user := range users {
		if user.Email == userLogin.Email && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)) == nil {
			if user.Ban == 0 {
				tokenStr := createToken(user)
				cookieOrSession(w, r, userLogin.SaveInfo, tokenStr)
				w.WriteHeader(http.StatusOK)
				return
			} else {
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func createToken(user User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["user-id"] = user.ID
	claim["user-role"] = user.Role
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenStr, err := token.SignedString([]byte("token-user"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenStr
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
