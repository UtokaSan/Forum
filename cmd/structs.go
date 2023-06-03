package cmd

type User struct {
	ID       int
	Image    string
	Username string
	Email    string
	Password string
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
	Username string `json:"username"`
	Password string `json:"password"`
}
