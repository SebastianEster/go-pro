package smtp

import "fmt"

type SenderImpl struct {}

func (sender *SenderImpl) SendEmail(address string, email string) {
	fmt.Printf("Sent email with content %s to address %s", email, address)
}

func NewSenderImpl() *SenderImpl {
	return new(SenderImpl)
}
