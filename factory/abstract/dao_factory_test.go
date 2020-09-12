package abstract

import "testing"

func TestMySQLFactory_CreateOrderMainDAO(t *testing.T) {
	var factory DAOFactory
	factory = new(MySQLFactory)
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().ExportOrderDetail()

	factory = new(OracleFactory)
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().ExportOrderDetail()
}