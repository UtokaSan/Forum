package cmd

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

func (u User) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

type Post struct {
	ID          int
	IDCreator   int
	NameCreator string
	Photo       string
	Title       string
	Texte       string
	Hidden      int
	Like        int
	Dislike     int
	Signalement int
	Categorie   string
	Ban         int
	Archived    string
}

type Comment struct {
	ID        int    `json:"ID"`
	IDPost    int    `json:"IDPost"`
	IDCreator int    `json:"IDCreator"`
	Text      string `json:"text"`
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
type Input struct {
	ID int `json:"id"`
}
type TakePostId struct {
	Info []map[string]interface{} `json:"info"`
}

type GestionPost struct {
	ID         int    `json:"id"`
	CreatePost string `json:"create-post"`
}

type AdminPanelChange struct {
	Key           string `json:"key"`
	UnBanUser     string `json:"unban-user"`
	BanUser       string `json:"ban-user"`
	RoleAdminUser string `json:"role-admin-user"`
	RoleModoUser  string `json:"role-modo-user"`
	DeletePost    string `json:"delete-post"`
}

type Register struct {
	Nom      string `json:"pseudo"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseRegister struct {
	Message string `json:"message"`
}

type responseLoginGithub struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	//Accept       string `json:"accept"`
}

type DataTokenJWT struct {
	UserId   int `json:"user-id"`
	UserRole int `json:"user-role"`
	Exp      int `json:"exp"`
}

type UserGoogle struct {
	Email         string `json:"email"`
	Nom           string `json:"name"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}
type Reaction struct {
	Reactions string `json:"reactions"`
	PostId    string `json:"post_id"`
}
