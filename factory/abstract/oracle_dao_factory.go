package abstract

import "fmt"

type OracleOrderMainDAO struct {
}

func (m *OracleOrderMainDAO) SaveOrderMain() {
	fmt.Println("oracle order main, save method")
}

type OracleOrderDetailDAO struct {
}

func (m *OracleOrderDetailDAO) ExportOrderDetail() {
	fmt.Println("oracle order detail, export method")
}
