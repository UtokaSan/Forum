package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"net/http"
	"time"
)

func loginPost(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var userLogin Login
	err = json.Unmarshal(body, &userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, user := range readUser() {
		if user.Email == userLogin.Email && userLogin.Password == user.Password {
			claim := token.Claims.(jwt.MapClaims)
			claim["sub"] = user.Username
			claim["exp"] = time.Now().Add(time.Hour * 24).Unix()
			tokenStr, err := token.SignedString([]byte("token-user"))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(tokenStr)
			r.Header.Set("Authorization", "Bearer "+tokenStr)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
