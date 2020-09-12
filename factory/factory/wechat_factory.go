package factory

import "fmt"

type WeChatAlertFactory struct {
	
}

func (w *WeChatAlertFactory) CreateAlert() Alert {
	return &WeChatAlert{}
}

type WeChatAlert struct {
	
}

func (w *WeChatAlert) Send(message string) {
	fmt.Println("wechat send success, message:", message)
}


