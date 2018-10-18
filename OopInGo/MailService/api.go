package MailService

type Sender interface {
	SendEmail(address string, email string)
}
