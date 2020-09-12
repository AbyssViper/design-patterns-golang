package simple

import "testing"

func TestNewAlert(t *testing.T) {
	alertManager := NewAlert("sms")
	alertManager.Send("send by sms")

	alertManager2 := NewAlert("mail")
	alertManager2.Send("send by mail")

	alertManager3 := NewAlert("wechat")
	alertManager3.Send("send by wechat")
}