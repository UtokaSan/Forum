package cmd

type User struct {
	ID       int
	Image    string
	Username string
	Email    string
	Password string
	Role     string
	Ban      int
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
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	SaveInfo string `json:"saveinfo"`
	JwtToken string `json:"jwtToken"`
}

type Register struct {
	Nom      string `json:"pseudo"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseRegister struct {
	Message string `json:"message"`
}
