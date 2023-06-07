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
}
