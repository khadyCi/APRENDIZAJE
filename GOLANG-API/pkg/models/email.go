package models

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	//CC      string `json:"cc"`
	Message string `json:"message"`
}
