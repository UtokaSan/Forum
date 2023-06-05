package cmd

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user Register
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Second part on program

	var userSend, errStr = changeRegisterToUser(user)

	if errStr != "" {
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

	//Create User

	createUser(userSend)
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

	fmt.Println("userDBEmail : " + userDBEmail + "| userDBUsername : " + userDBUsername)

	if userDBUsername == "" && userDBEmail == "" {
		return false
	}
	return true
}

func changeRegisterToUser(user Register) (User, string) {
	if strings.Contains(strings.ToUpper(user.Nom), strings.ToUpper("Jordan")) && strings.Contains(strings.ToUpper(user.Email), strings.ToUpper("Jordan")) {
		return User{}, "OOHH no, Sorry you can't create a Account 😉"
	}

	var userSend User

	userSend.Email = user.Email
	userSend.Password = cryptPassword(user.Password)
	userSend.Username = user.Nom

	return userSend, ""
}
