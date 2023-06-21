package cmd

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v53/github"
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
	if errroCreateUserEmpty(userSend) {
		createErrorMessage("Un champ de texte est vide", 403, w)
		return
	}

	userSend = createUser(userSend)
	createAToken(w, r, userSend)

	_, err = w.Write(createSuccessfulMessage("compte bien créer", 201, w))
	if err != nil {
		return
	}
}

func CreateAccountGoogle(r *http.Request) (User, bool) {
	account := getInfoGoogle(r)
	account.Nom = strings.Split(account.Email, "@")[0]
	if account.VerifiedEmail == true {
		user := readOneUserByIdentifiantWithGoogle(account.Email)
		if user.ID == -1 {
			if checkInputNotValid(account.Email, account.Nom) {
				return User{}, true
			}
			createUserGoogle(account)
			return User{}, true
		}
		return user, false
	}
	return User{}, true

}

func CreateUserGithub(w http.ResponseWriter, r *http.Request, user User) (User, bool) {

	userDB := readOneUserByIdentifiantWithGoogle(user.Email)

	if userDB.ID == -1 {
		if checkInputNotValid(user.Email, user.Username) {

		}
		user = createUserGithub(user)
		tokenStr := createToken(user)
		var store = sessions.NewCookieStore([]byte("secret-key"))
		session, err := store.Get(r, "session-login")
		if err != nil {
			fmt.Println(err)
		}
		session.Values["jwtToken"] = tokenStr
		err = session.Save(r, w)
		return user, false
	}

	tokenStr := createToken(userDB)
	var store = sessions.NewCookieStore([]byte("secret-key"))
	session, err := store.Get(r, "session-login")
	if err != nil {
		fmt.Println(err)
	}
	session.Values["jwtToken"] = tokenStr
	err = session.Save(r, w)

	return userDB, true
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
	fmt.Println(user.ID)

	claim["user-id"] = user.ID
	claim["user-role"] = user.Role
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenStr, err := token.SignedString([]byte("token-user"))
	fmt.Println(tokenStr)
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

func updateUnBanUserOrBan(user User) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := "UPDATE users SET ban = ? WHERE nom = ?"
	_, err = db.Exec(query, user.Ban, user.Username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User unban update successfully")
}

func updateUserRole(user User) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	query := "UPDATE users SET role = ? WHERE nom = ?"
	_, err = db.Exec(query, user.Role, user.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User role updated successfully")
}

func convGithubToUser(userGithub *github.User, client *github.Client) User {
	ctx := context.Background()
	emails, _, _ := client.Users.ListEmails(ctx, nil)

	email := ""

	fmt.Println("url : ", userGithub.GetAvatarURL())

	if len(emails) > 1 {
		email = *emails[1].Email
	} else {
		fmt.Println("Aucune adresse email trouvée.")
		return User{ID: -1}
	}

	return User{
		Image:    userGithub.GetAvatarURL(),
		Username: userGithub.GetName(),
		Email:    email,
		Role:     "1",
	}
}

func errroCreateUserEmpty(user User) bool {
	if user.Password == "" || user.Email == "" || user.Username == "" {
		return true
	}
	return false
}
