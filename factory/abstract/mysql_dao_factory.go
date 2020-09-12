package abstract

import "fmt"

type MySQLOrderMainDAO struct {
}

func (m *MySQLOrderMainDAO) SaveOrderMain() {
	fmt.Println("mysql order main, save method")
}

type MySQLOrderDetailDAO struct {
}

func (m *MySQLOrderDetailDAO) ExportOrderDetail() {
	fmt.Println("mysql order detail, export method")
}
