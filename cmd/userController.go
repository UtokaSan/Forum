package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user Register
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("mince -1")

	if err != nil {
		fmt.Println("mince 0")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var userSend, errStr = changeRegisterToUser(user)

	if errStr != "" {
		println("mince 1")
		w.WriteHeader(http.StatusUnauthorized)
		_, err := fmt.Fprintln(w, err)
		if err != nil {
			return
		}
	}

	if userAlreadyExist(userSend) {
		fmt.Println("Error Account already Create please change the email or password")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintln(w, "Error Account already Create please change the email or password")
		if err != nil {
			return
		}
	}

	fmt.Println("mince 999")

	//createUser(userSend)
	createAToken(w, r, userSend)
	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintln(w, "creation of account successful")
	if err != nil {
		return
	}
}

func cryptPassword(password string) string {
	passwordCrypt, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println("error with crypt")
	}
	return string(passwordCrypt)
}

func userAlreadyExist(user User) bool {

	userDBEmail := readOneUserByEmailOrPseudo(user.Email).Email
	userDBUsername := readOneUserByEmailOrPseudo(user.Username).Username

	//fmt.Println("userDBEmail : " + userDBEmail + "| userDBUsername : " + userDBUsername)

	if userDBUsername == "" && userDBEmail == "" {
		return false
	}
	return true
}

func changeRegisterToUser(user Register) (User, string) {
	if strings.Contains(strings.ToUpper(user.Nom), strings.ToUpper("Jordan")) && strings.Contains(strings.ToUpper(user.Email), strings.ToUpper("Jordan")) {
		fmt.Println("MERDE")
		return User{}, "OOHH no, Sorry you can't create a Account 😉"
	}

	var userSend User

	userSend.Email = user.Email
	userSend.Password = cryptPassword(user.Password)
	userSend.Username = user.Nom

	fmt.Println(userSend)
	fmt.Println("user.Email : " + user.Email)
	fmt.Println("user.Nom : " + user.Nom)
	fmt.Println(user.Password)

	return userSend, ""
}

func createAToken(w http.ResponseWriter, r *http.Request, user User) {
	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["user-id"] = user.ID
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenStr, err := token.SignedString([]byte("token-user"))
	if err != nil {
		fmt.Println(err)
		return
	}

	var store = sessions.NewCookieStore([]byte("secret-key"))
	session, err := store.Get(r, "session-login")
	if err != nil {
		fmt.Println(err)
	}
	session.Values["jwtToken"] = tokenStr
	err = session.Save(r, w)

	return
}
