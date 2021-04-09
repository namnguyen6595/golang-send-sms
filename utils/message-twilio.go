package utils

import (
	"encoding/json"
	"net/http"
	"net/url"
	_ "project/const"
	_const "project/const"
	"strings"
)


type Response struct {
	Status int64 `json:"status"`
	Message string `json:"message"`
}



const accountSid = _const.AccountSid
const authToken = _const.AuthToken

const urlService = "https://api.twilio.com/2010-04-01/Accounts/"+accountSid+"/Messages.json"

func ServicesSMS(to []string, body string) (Response) {
	urlStr := urlService
	bodyMessage := body

	var status int64
	var message string

	msgData := url.Values{}
	msgData.Set("From", "+12532521307")

	for _, numb := range to {
		msgData.Set("To", "+84"+numb)
		msgData.Set("Body", bodyMessage)
		msgDataReader := *strings.NewReader(msgData.Encode())


		client := &http.Client{}
		req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
		req.SetBasicAuth(accountSid, authToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		response,_ := client.Do(req)
		if response.StatusCode >= 200 && response.StatusCode < 300 {
			var data map[string]interface{}

			decoder := json.NewDecoder(response.Body)
			err := decoder.Decode(&data)
			if err == nil {
				status = http.StatusOK
				message = "Send success"
			}
		} else {
			status = http.StatusBadRequest
			message = "Phonemes invalid"
		}
	}
	return Response{status, message}
}