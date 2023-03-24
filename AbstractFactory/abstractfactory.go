package AbstractFactory

import "fmt"

// ProductMainDAO 为产品主记录
type ProductMainDAO interface {
	SaveProductMain()
}

// ProductDetailDAO 为产品详细记录
type ProductDetailDAO interface {
	SaveProductDetail()
}

// DAOFactory 抽象模式工厂接口
type DAOFactory interface {
	CreateProductMainDAO() ProductMainDAO
	CreateProductDetailDAO() ProductDetailDAO
}

// RDBMainDAO 关系型数据库的Main实现
type RDBMainDAO struct {
}

func (r *RDBMainDAO) SaveProductMain() {
	fmt.Print("rdb main save\n")
}

// RDBDetailDAO 关系型数据库的Detail实现
type RDBDetailDAO struct {
}

func (r *RDBDetailDAO) SaveProductDetail() {
	fmt.Print("rdb detail save\n")
}

// RDBDAOFactory RDB的抽象工厂实现
type RDBDAOFactory struct {
}

func (r *RDBDAOFactory) CreateProductMainDAO() ProductMainDAO {
	return &RDBMainDAO{}
}

func (r *RDBDAOFactory) CreateProductDetailDAO() ProductDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO XML存储的Main实现
type XMLMainDAO struct {
}

func (x *XMLMainDAO) SaveProductMain() {
	fmt.Print("xml main save\n")
}

// XMLDetailDAO XML存储的Detail实现
type XMLDetailDAO struct {
}

func (x *XMLDetailDAO) SaveProductDetail() {
	fmt.Print("xml detail save\n")
}

// XMLDAOFactory XML抽象工厂的实现
type XMLDAOFactory struct {
}

func (x *XMLDAOFactory) CreateProductMainDAO() ProductMainDAO {
	return &XMLMainDAO{}
}

func (x *XMLDAOFactory) CreateProductDetailDAO() ProductDetailDAO {
	return &XMLDetailDAO{}
}
