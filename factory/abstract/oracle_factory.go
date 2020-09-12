package abstract


type OracleFactory struct {
}

func (m *OracleFactory) CreateOrderMainDAO() OrderMainDAO {
	return &OracleOrderMainDAO{}
}

func (m *OracleFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &OracleOrderDetailDAO{}
}

