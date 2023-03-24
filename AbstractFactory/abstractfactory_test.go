package AbstractFactory

import "testing"

func getMainAndDetailFactory(factory DAOFactory) {
	factory.CreateProductMainDAO().SaveProductMain()
	factory.CreateProductDetailDAO().SaveProductDetail()
}

func TestRDBFactory(t *testing.T) {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetailFactory(factory)
}

func TestXMLFactory(t *testing.T) {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetailFactory(factory)
}
