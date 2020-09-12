package factory

import "testing"

func TestMailAlert_Send(t *testing.T) {
	alertFactory := new(SMSAlertFactory)
	alert := alertFactory.CreateAlert()
	alert.Send("hello")
}