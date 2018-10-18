package client

import (
	"github.com/SebastianEster/go-pro/OopInGo/MailService"
	"github.com/SebastianEster/go-pro/OopInGo/MailService/util"
)

func SendEmail(address, email string) {
	sender := util.Get("mail.sender").(MailService.Sender)
	sender.SendEmail(address, email)
}
