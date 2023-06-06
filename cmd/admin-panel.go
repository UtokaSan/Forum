package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func adminPanel(w http.ResponseWriter, r *http.Request) {
	userPanel := AdminPanel{
		Account:      takeAllUsers(),
		Ban:          takeUsersBan(),
		PostHidden:   takePostHidden(),
		PostArchived: postArchived(),
	}
	jsonData, err := json.Marshal(userPanel)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(jsonData)
}
