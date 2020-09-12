package factory

// 工厂接口
type AlertFactory interface {
	CreateAlert() Alert
}


type Alert interface {
	Send(message string)
}
