package FactoryMethod

// Product 一个抽象的产品
type Product interface {
	SetName(name string)
	GetName() string
}

// Product1 实现具体的产品：产品1
type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return "产品1的name为" + p1.name
}

// Product2 实现具体的产品：产品2
type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return "产品2的name为" + p2.name
}

// ProductFactory 实现一个抽象工厂
type ProductFactory interface {
	Create() Product
}

// Product1Factory 实现产品1的工厂
type Product1Factory struct {
}

func (p1f *Product1Factory) Create() Product {
	return &Product1{}
}

// 实现产品2的工厂
type Product2Factory struct {
}

func (p2f *Product2Factory) Create() Product {
	return &Product2{}
}