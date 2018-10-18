package client

import (
	"github.com/SebastianEster/go-pro/OopInGo/MailService/smtp"
	"github.com/SebastianEster/go-pro/OopInGo/MailService/util"
	"testing"
)

func init() {
	util.Register("mail.sender", smtp.NewSenderImpl())
}

func TestSendEmail(t *testing.T) {
	SendEmail("sebastian.ester.se@gmail.com", "Hello World!")
}
