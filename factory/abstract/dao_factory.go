package abstract

type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type OrderMainDAO interface {
	SaveOrderMain() // 存储主订单
	// ...其他的方法
}

type OrderDetailDAO interface {
	ExportOrderDetail() // 导出订单详情
	// ...其他的方法
}
