package factory

import "fmt"

type Alert interface {
	Send(message string)
}

func NewAlert(alertType string) Alert {
	if alertType == "mail" {
		return &Mail{}
	} else if alertType == "wechat" {
		return &WeChat{}
	// 横向拓展后，通过添加到生产入口并增加标识，即可
	} else if alertType == "sms" {
		return &SMS{}
	} else {
		return nil
	}
}

type Mail struct {
}

func (m *Mail) Send(message string) {
	fmt.Println("mail send success, message:", message)
}

type WeChat struct {
}

func (w *WeChat) Send(message string) {
	fmt.Println("wechat send success, message:", message)
}

// 如果拓展短信业务，增加相关结构体，并实现相关接口
type SMS struct {

}

func (S *SMS) Send(message string) {
	fmt.Println("sms send success, message:", message)
}
