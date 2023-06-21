package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

func authGuestSecurity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := getSession(r)
		fmt.Println()
		fmt.Println("TOKEN : ", token)
		fmt.Println()
		if token == "" {
			tokenCookie := getCookie(r)
			fmt.Println("MEdedededeD : ", tokenCookie)

			if tokenCookie == "" {
				next(w, r)
				return
			}

			var store = sessions.NewCookieStore([]byte("secret-key"))
			session, _ := store.Get(r, "session-login")
			session.Values["jwtToken"] = tokenCookie
			session.Save(r, w)
			fmt.Println("session : ", session)
			next(w, r)
			return
		}
		next(w, r)
		return
	}
}

func authUserSecurity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var secretToken = "token-user"
		token := getSession(r)

		println("token : " + token)

		if token == "" {
			tokenCookie := getCookie(r)
			if tokenCookie == "" {
				next(w, r)
				return
			}

			var store = sessions.NewCookieStore([]byte("secret-key"))
			session, _ := store.Get(r, "session-login")
			session.Values["jwtToken"] = tokenCookie
			session.Save(r, w)
			fmt.Println("session : ", session)
		}
		tokenJWT := checkJWT("", token)
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
			tokenCookie := getCookie(r)
			fmt.Println("Coooookie : ", tokenCookie)
			if tokenCookie == "" {
				next(w, r)
				return
			}
			var store = sessions.NewCookieStore([]byte("secret-key"))
			session, err := store.Get(r, "session-login")
			if err != nil {
				fmt.Println(err)
			}
			session.Values["jwtToken"] = tokenCookie
			err = session.Save(r, w)
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
			tokenCookie := getCookie(r)
			if tokenCookie == "" {
				next(w, r)
				return
			}
			var store = sessions.NewCookieStore([]byte("secret-key"))
			session, _ := store.Get(r, "session-login")
			session.Values["jwtToken"] = tokenCookie
			session.Save(r, w)
			fmt.Println("session : ", session)
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
	fmt.Println("MERDEDEEEEEEEEFCDSxw : ", session)

	if session.Values["jwtToken"] == nil {
		return ""
	}
	return session.Values["jwtToken"].(string)
}

func getCookie(r *http.Request) string {
	cookieUser, err := r.Cookie("jwtToken")
	if err != nil {
		println("pas de cookie")
		return ""
	}
	cookieStr := cookieUser.String()
	return cookieStr[9:len(cookieStr)]
}

func checkJWT(secretToken string, tokenJWT string) *jwt.Token {
	token, err := jwt.Parse(tokenJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Non")
			return nil, fmt.Errorf("Méthode de signature inattendue : %v", token.Header["alg"])
		}
		fmt.Println("Yes")
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
		fmt.Println("mince ! avec le dataTokenJWT qui est dans checkJWT l:152")
		return DataTokenJWT{
			UserRole: 0,
		}
	}
	allDataToken := token.Claims.(jwt.MapClaims)

	fmt.Println("----------------------------------")
	fmt.Println(allDataToken["user-id"])
	fmt.Println(allDataToken["user-role"])
	fmt.Println(allDataToken["exp"])
	fmt.Println("----------------------------------")

	// Accéder aux données du JWT

	data.UserId = int(allDataToken["user-id"].(float64))
	data.UserRole, _ = strconv.Atoi(allDataToken["user-role"].(string))
	data.Exp = int(allDataToken["exp"].(float64))

	fmt.Println("user-role:", data)
	return data
}
