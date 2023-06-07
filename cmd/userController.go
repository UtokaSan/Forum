package cmd

import (
	"database/sql"
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

	if err != nil {
		createErrorMessage("bug with request", 500, w)
		return
	}

	var userSend, errStr = changeRegisterToUser(user)
	if errStr != "" {
		createErrorMessage(errStr, 403, w)
		return
	}

	if userAlreadyExist(userSend) {
		createErrorMessage("Compte déjà existant", 403, w)
		return
	}

	createUser(userSend)
	createAToken(w, r, userSend)

	_, err = w.Write(createSuccessfulMessage("compte bien créer", 201, w))
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

	if userDBUsername == "" && userDBEmail == "" {
		return false
	}
	return true
}

func changeRegisterToUser(user Register) (User, string) {
	if strings.Contains(strings.ToUpper(user.Nom), strings.ToUpper("Jordan")) || strings.Contains(strings.ToUpper(user.Email), strings.ToUpper("Jordan")) {
		return User{}, "OOHH no, Sorry you can't create a Account 😉"
	}

	var userSend User

	userSend.Email = user.Email
	userSend.Password = cryptPassword(user.Password)
	userSend.Username = user.Nom

	return userSend, ""
}

func createAToken(w http.ResponseWriter, r *http.Request, user User) {
	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["user-id"] = user.ID
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenStr, err := token.SignedString([]byte("token-user"))
	if err != nil {
		createErrorMessage("Bug avec le token d'authentification", 500, w)
		return
	}

	var store = sessions.NewCookieStore([]byte("secret-key"))
	session, err := store.Get(r, "session-login")
	if err != nil {
		createErrorMessage("Bug avec le token authentication", 500, w)
		return
	}
	session.Values["jwtToken"] = tokenStr
	err = session.Save(r, w)

	return
}

func createErrorMessage(message string, code int, w http.ResponseWriter) {
	fmt.Println("mince : " + message)

	errorMessage := responseRegister{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonData, err := json.Marshal(errorMessage)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

func createSuccessfulMessage(message string, code int, w http.ResponseWriter) []byte {
	fmt.Println("successfully : " + message)

	errorMessage := responseRegister{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonData, err := json.Marshal(errorMessage)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return nil
	}

	return jsonData
}

func updateUserBan(user User) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "UPDATE users SET  nom = ?, ban = ? WHERE nom = ?"
	_, err = db.Exec(query, user.Username, user.Ban, user.Username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User ban update successfully")
}
