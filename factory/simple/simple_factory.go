package simple

import "fmt"

type Alert interface {
	Send(message string)
}

func NewAlert(alertType string) Alert {
	switch alertType {
	case "mail":
		return &Mail{}
	case "wechat":
		return &WeChat{}
	case "sms":
		return &SMS{}
	default:
		return nil
	}
}

// Product Mail
type Mail struct {
}

func (m *Mail) Send(message string) {
	fmt.Println("mail send success, message:", message)
}

// Product WeChat
type WeChat struct {
}

func (w *WeChat) Send(message string) {
	fmt.Println("wechat send success, message:", message)
}

// Product SMS
// 如果拓展短信业务，增加相关结构体，并实现相关接口
type SMS struct {
}

func (S *SMS) Send(message string) {
	fmt.Println("sms send success, message:", message)
}
