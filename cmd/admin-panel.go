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
		Ban:             takeUsersBan(),
		PostHidden:      takePostHidden(),
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
	fmt.Println(string(body[2:12]))
	if string(body[2:12]) == "deban-user" {
		updateUserBan(User{
			Username: change.DebanUser,
			Ban:      0,
		})
	} else if string(body[2:11]) == "user-role" {
		updateUser(User{
			Role: "admin",
		})
	} else {
		fmt.Println("Nique ta mere jordan")
	}
}
