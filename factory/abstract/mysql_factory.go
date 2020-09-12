package abstract


type MySQLFactory struct {
}

func (m *MySQLFactory) CreateOrderMainDAO() OrderMainDAO {
	return &MySQLOrderMainDAO{}
}

func (m *MySQLFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &MySQLOrderDetailDAO{}
}

