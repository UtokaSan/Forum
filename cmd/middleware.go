package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

func authUserSecurity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const secretToken = "token-user"

		token := getSession(r)
		if token == "" {
			http.Redirect(w, r, "/login", 401)
			return
		}
		tokenJWT := checkJWT(secretToken, token)
		dataUser := getData(tokenJWT)

		if dataUser.UserRole >= 1 {
			fmt.Println("user : ", dataUser.UserRole)
			next(w, r)
		} else {
			http.Redirect(w, r, "/login", 401)
			return
		}
	}
}

func authModoSecurity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const secretToken = "token-user"

		token := getSession(r)
		fmt.Println("token : " + token)
		//fmt.Println("token : ", len(token))
		if token == "" {
			http.Redirect(w, r, "/login", 401)
			return
		}
		tokenJWT := checkJWT(secretToken, token)
		dataUser := getData(tokenJWT)

		if dataUser.UserRole >= 2 {
			fmt.Println("user : ", dataUser.UserRole)
			next(w, r)
		} else {
			http.Redirect(w, r, "/login", 401)
			return
		}
	}
}

func authAdminSecurity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const secretToken = "token-user"

		token := getSession(r)
		fmt.Println("token : " + token)
		//fmt.Println("token : ", len(token))
		if token == "" {
			http.Redirect(w, r, "/login", 401)
			return
		}
		tokenJWT := checkJWT(secretToken, token)
		dataUser := getData(tokenJWT)

		if dataUser.UserRole >= 3 {
			fmt.Println("user : ", dataUser.UserRole)
			next(w, r)
		} else {
			fmt.Println("user : ", dataUser.UserRole)
			http.Redirect(w, r, "/login", 401)
			return
		}
	}
}

// ----------------------------------------------------
// |                    OTHER                         |
// ----------------------------------------------------

func getSession(r *http.Request) string {
	var store = sessions.NewCookieStore([]byte("secret-key"))
	session, _ := store.Get(r, "session-login")
	if session.Values["jwtToken"] == nil {
		return ""
	}
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

func getData(token *jwt.Token) DataTokenJWT {
	data := DataTokenJWT{}
	if token == nil {
		return DataTokenJWT{
			UserRole: 0,
		}
	}
	allDataToken := token.Claims.(jwt.MapClaims)

	fmt.Println("----------------------------------")
	fmt.Println(allDataToken["user-id"])
	fmt.Println(allDataToken["user-id"].(float64))
	fmt.Println(allDataToken["user-role"])
	fmt.Println(strconv.Atoi(allDataToken["user-role"].(string)))
	fmt.Println(allDataToken["exp"])
	fmt.Println(allDataToken["exp"].(float64))
	fmt.Println("----------------------------------")

	// Accéder aux données du JWT
	data.UserId = allDataToken["user-id"].(float64)
	data.UserRole, _ = strconv.Atoi(allDataToken["user-role"].(string))
	//data.Exp = allDataToken["user-fdsfdsqf"].(float64)

	fmt.Println("user-role:", data)
	return data
}
