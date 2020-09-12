package factory

import "fmt"

type SMSAlertFactory struct {

}

func (w *SMSAlertFactory) CreateAlert() Alert {
	return &SMSAlert{}
}

type SMSAlert struct {

}

func (w *SMSAlert) Send(message string) {
	fmt.Println("sms send success, message:", message)
}



