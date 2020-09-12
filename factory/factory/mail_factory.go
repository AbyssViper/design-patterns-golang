package factory

import "fmt"

type MailAlertFactory struct {

}

func (m *MailAlertFactory) CreateAlert() Alert {
	return &MailAlert{}
}

type MailAlert struct {

}

func (m *MailAlert) Send(message string) {
	fmt.Println("mail send success, message:", message)
}
