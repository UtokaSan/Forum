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
	"strconv"
	"strings"
	"time"
)

func loginGoogle(w http.ResponseWriter, r *http.Request) {
	config := getConfig()
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func loginGithub(w http.ResponseWriter, r *http.Request) {
	CLIENT_ID := "2289380b3bb541be1a3a"
	//CLIENT_SECRET := "81443484b632c86271768c67ee7a4da0c6e8ee0e "

	http.Redirect(w, r, "https://github.com/login/oauth/authorize?scope=user:email&client_id="+CLIENT_ID, http.StatusTemporaryRedirect)
}

func callbackLoginGoogle(w http.ResponseWriter, r *http.Request) {
	account, err := CreateAccountGoogle(r)
	if err {
		return
	}
	fmt.Println("account ", account)
	//createAToken(w, r, account)

	userRole, errAtoi := strconv.Atoi(account.Role)
	fmt.Println("userRole : ", userRole)
	fmt.Println("err : ", errAtoi)
	if userRole == 1 || userRole == 2 {
		http.Redirect(w, r, "http://localhost:8080/login", http.StatusTemporaryRedirect) // Pour le moment
	} else if userRole == 3 {
		http.Redirect(w, r, "http://localhost:8080/register", http.StatusTemporaryRedirect) // Pour le moment
	}

}

func callbackLoginGithub(w http.ResponseWriter, r *http.Request) {

	//session_code := request.env['rack.request.query_hash']['code']
	code := r.URL.Query()
	fmt.Println(code)
	fmt.Println("-----------")
	fmt.Println(code.Get("code"))

	//CLIENT_ID := "2289380b3bb541be1a3a"
	//CLIENT_SECRET := "81443484b632c86271768c67ee7a4da0c6e8ee0e "

	//# ... and POST it back to GitHub

	//result = RestClient.post('https://github.com/login/oauth/access_token')

	//# extract the token and granted scopes
	//access_token = JSON.parse(result)['access_token']

	//account, err := CreateAccountGoogle(r)
	//if err {
	//	return
	//}
	//fmt.Println("account ", account)
	////createAToken(w, r, account)
	//
	//userRole, errAtoi := strconv.Atoi(account.Role)
	//fmt.Println("userRole : ", userRole)
	//fmt.Println("err : ", errAtoi)
	//if userRole == 1 || userRole == 2 {
	//	http.Redirect(w, r, "http://localhost:8080/login", http.StatusTemporaryRedirect) // Pour le moment
	//} else if userRole == 3 {
	//	http.Redirect(w, r, "http://localhost:8080/register", http.StatusTemporaryRedirect) // Pour le moment
	//}

}

func getConfig() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     "116188844729-bpmpofo72u5vdhdt43qif41lmppqejuh.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/admin",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
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
	token := jwt.New(jwt.SigningMethodHS256)
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
				claim := token.Claims.(jwt.MapClaims)
				claim["user-id"] = user.ID
				claim["user-role"] = user.Role
				claim["exp"] = time.Now().Add(time.Hour * 24).Unix()
				tokenStr, err := token.SignedString([]byte("token-user"))
				if err != nil {
					fmt.Println(err)
					return
				}
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

func CreateAccountGoogle(r *http.Request) (User, bool) {
	account := getInfoGoogle(r)
	account.Nom = strings.Split(account.Email, "@")[0]
	if account.VerifiedEmail == true {
		user := readOneUserByIdentifiantWithGoogle(account.Email)

		fmt.Println("--------------")
		fmt.Println("User2 : ", user)
		fmt.Println("--------------")
		if user.ID == -1 {
			if checkInputNotValid(account.Email, account.Nom) {
				return User{}, true
			}
			createUserGoogle(account)
			return convUserGoogleToUser(account), false
		}
		return user, false
	}
	return User{}, true

}

func getInfoGoogle(r *http.Request) UserGoogle {
	config := getConfig()
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

	fmt.Println("------)")
	fmt.Println(usergoogle)
	fmt.Println("------)")
	return usergoogle

}

func checkInputNotValid(email string, pseudo string) bool {
	if email == "" && pseudo == "" {
		return true
	}
	fmt.Println("email : ", email, ", pseudo : ", pseudo)
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
