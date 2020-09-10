package factory

import "testing"

func TestNewAlert(t *testing.T) {
	alertManager := NewAlert("sms")
	alertManager.Send("send by sms")

	alertManager2 := NewAlert("mail")
	alertManager2.Send("send by mail")
}