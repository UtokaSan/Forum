package cmd

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       int
	Image    string
	Username string
	Email    string
	Password string
	Role     string
	Ban      int
	Report   string
}

type Post struct {
	ID          int
	Photo       string
	Texte       string
	Hidden      int
	Like        int
	Dislike     int
	Signalement int
	Categorie   string
	Ban         int
	Archived    string
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	SaveInfo string `json:"saveinfo"`
	JwtToken string `json:"jwtToken"`
}

type AdminPanel struct {
	Account         []map[string]interface{} `json:"account"`
	AccountReported []map[string]interface{} `json:"accountReported"`
	Ban             []map[string]interface{} `json:"ban"`
	PostHidden      []map[string]interface{} `json:"postHidden"`
	PostArchived    []map[string]interface{} `json:"postArchived"`
}

type AdminPanelChange struct {
	Key           string `json:"key"`
	DebanUser     string `json:"deban-user"`
	RoleAdminUser string `json:"role-admin-user"`
	RoleModoUser  string `json:"role-modo-user"`
}

type Register struct {
	Nom      string `json:"pseudo"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseRegister struct {
	Message string `json:"message"`
}

type Claims struct {
	Userid int `json:"user-id"`
	Exp    int `json:"exp"`
}

func (c Claims) GetExpirationTime() (*jwt.NumericDate, error) {
	//TODO implement me
	panic("implement me")
}

func (c Claims) GetIssuedAt() (*jwt.NumericDate, error) {
	//TODO implement me
	panic("implement me")
}

func (c Claims) GetNotBefore() (*jwt.NumericDate, error) {
	//TODO implement me
	panic("implement me")
}

func (c Claims) GetIssuer() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c Claims) GetSubject() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c Claims) GetAudience() (jwt.ClaimStrings, error) {
	//TODO implement me
	panic("implement me")
}
