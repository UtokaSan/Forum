package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v53/github"
	_ "github.com/google/go-github/v53/github"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func loginGoogle(w http.ResponseWriter, r *http.Request) {
	config := getConfig("116188844729-bpmpofo72u5vdhdt43qif41lmppqejuh.apps.googleusercontent.com", "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", []string{"https://www.googleapis.com/auth/userinfo.email"}, google.Endpoint)

	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callbackLoginGoogle(w http.ResponseWriter, r *http.Request) {
	account, err := CreateAccountGoogle(r)
	if err {
		return
	}
	fmt.Println("account ", account)
	//createAToken(w, r, account)
}

func callbackLoginGithub(w http.ResponseWriter, r *http.Request) {
	user := getUserGithub(r)
	if user.ID == -1 {
		fmt.Println("error with Get User")
	}
	CreateUserGithub(w, r, user)

}

func loginGithub(w http.ResponseWriter, r *http.Request) {
	endpoint := oauth2.Endpoint{
		TokenURL: "https://github.com/login/oauth/access_token",
		AuthURL:  "https://github.com/login/oauth/authorize",
	}

	config := getConfig("2289380b3bb541be1a3a", "81443484b632c86271768c67ee7a4da0c6e8ee0e", []string{"user:email"}, endpoint)
	url := config.AuthCodeURL("state", oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func getConfig(clientID string, clientSecret string, auth []string, endpoint oauth2.Endpoint) *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/api/callbacklogingithub",
		Scopes:       auth,
		Endpoint:     endpoint,
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

func getInfoGoogle(r *http.Request) UserGoogle {
	fmt.Println("--------------")
	config := getConfig("116188844729-bpmpofo72u5vdhdt43qif41lmppqejuh.apps.googleusercontent.com", "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", []string{"https://www.googleapis.com/auth/userinfo.email"}, google.Endpoint)
	code := r.URL.Query().Get("code")
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println(err)
		return UserGoogle{}
	}
	client := config.Client(oauth2.NoContext, token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		fmt.Println(err)
		return UserGoogle{}
	}
	body, err := ioutil.ReadAll(response.Body)

	var usergoogle UserGoogle
	err = json.Unmarshal(body, &usergoogle)
	if err != nil {
		fmt.Println(err)
		return UserGoogle{}
	}
	return usergoogle
}

func getInfoGithub(r *http.Request) UserGoogle {
	fmt.Println("--------------")
	config := getConfig("116188844729-bpmpofo72u5vdhdt43qif41lmppqejuh.apps.googleusercontent.com", "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", []string{"https://www.googleapis.com/auth/userinfo.email"}, google.Endpoint)
	code := r.URL.Query().Get("code")
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println(err)
		return UserGoogle{}
	}
	fmt.Println(token)
	client := config.Client(oauth2.NoContext, token)

	//ctx := context.Background()
	//githubClient := github.NewClient(client)

	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		fmt.Println(err)
		return UserGoogle{}
	}
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println("---------------")
	fmt.Println(body)

	var usergoogle UserGoogle
	err = json.Unmarshal(body, &usergoogle)
	if err != nil {
		fmt.Println(err)
		return UserGoogle{}
	}

	fmt.Println("------)")
	fmt.Println(usergoogle)
	fmt.Println("------)")
	return usergoogle
}

func checkInputNotValid(email string, pseudo string) bool {
	if email == "" && pseudo == "" {
		return true
	}
	return false
}

func convUserGoogleToUser(userGoogle UserGoogle) User {
	return User{
		Image:    userGoogle.Picture,
		Email:    userGoogle.Email,
		Username: userGoogle.Nom,
		Role:     "1",
	}
}

func getUserGithub(r *http.Request) User {
	code := r.URL.Query().Get("code")

	endpoint := oauth2.Endpoint{
		TokenURL: "https://github.com/login/oauth/access_token",
		AuthURL:  "https://github.com/login/oauth/authorize",
	}

	config := getConfig("2289380b3bb541be1a3a", "81443484b632c86271768c67ee7a4da0c6e8ee0e", []string{"user:email"}, endpoint)

	ctx := context.Background()
	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)
	client := github.NewClient(oauth2.NewClient(ctx, tokenSource))

	UserGithub, _, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Println("error to get user github :", err)
	}
	return convGithubToUser(UserGithub, client)
}
