package cmd

type User struct {
	ID       int
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
