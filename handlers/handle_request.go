package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type BodyMessage struct {
	To string
	Message string
}

func BodyRequest (w http.ResponseWriter, r *http.Request) {
	var msg BodyMessage
	err := json.NewDecoder(r.Body).Decode(&msg)
	numbers := strings.Split(msg.To, ",")
	resp :=SendSMSToEveryOne(numbers, msg.Message)
		 if err != nil {
		 	//resp, _ := json.Marshal(err)
		 	http.Error(w, err.Error(), http.StatusBadRequest)
		 }
		 fmt.Fprint(w, "Body", resp)

}

