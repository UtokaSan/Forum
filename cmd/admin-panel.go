package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendInfoAdmin(w http.ResponseWriter, r *http.Request) {
	userPanel := AdminPanel{
		Account:         takeAllUsers(),
		AccountReported: takeUserReported(),
		PostHidden:      takePostHidden(),
		PostUnHidden:    takePostUnHidden(),
		PostArchived:    postArchived(),
	}
	jsonData, err := json.Marshal(userPanel)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(jsonData)
}

func adminPanel(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var change AdminPanelChange
	err = json.Unmarshal(body, &change)
	switch change.Key {
	case "ban-user":
		updateUnBanUserOrBan(User{
			Username: change.BanUser,
			Ban:      1,
		})
	case "unban-user":
		updateUnBanUserOrBan(User{
			Username: change.UnBanUser,
			Ban:      0,
		})
	case "role-admin-user":
		updateUserRole(User{
			Username: change.RoleAdminUser,
			Role:     "admin",
		})
	case "role-modo-user":
		updateUserRole(User{
			Username: change.RoleModoUser,
			Role:     "modo",
		})
	case "delete-post":
		deletePost(Post{
			Title: change.DeletePost,
		})
	case "hidden-post":
		updateHiddenPost(Post{
			Title:  change.HiddenPost,
			Hidden: 0,
		})
	}
}
