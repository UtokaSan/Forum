package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"net/http"
)

func authUserSecurity(w http.ResponseWriter, r *http.Request) {
	const secretToken = "token-user"
	//tokenJWT := getSession(r)
	//tokenUnscript := checkJWT(secretToken, tokenJWT)

}

func getSession(r *http.Request) string {
	var store = sessions.NewCookieStore([]byte("secret-key"))
	session, _ := store.Get(r, "session-login")

	return session.Values["jwtToken"].(string)
}

func checkJWT(secretToken string, tokenJWT string) *jwt.Token {
	token, err := jwt.Parse(tokenJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Méthode de signature inattendue : %v", token.Header["alg"])
		}
		return []byte(secretToken), nil
	})
	if err != nil || !token.Valid {
		return nil
	}
	if token.Valid {
		_, err := token.Claims.(jwt.MapClaims)
		if err {
			fmt.Println("bug : ", err)
		}
	}
	return token
}
