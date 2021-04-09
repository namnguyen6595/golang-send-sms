package handlers

import (
	"project/utils"
)

type Data struct {
	To []string
	Message string
}

func SendSMSToEveryOne(to []string, body string) utils.Response {
	data := utils.ServicesSMS(to, body)

	return data
}
