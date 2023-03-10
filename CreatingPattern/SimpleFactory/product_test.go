package SimpleFactory

import (
	"fmt"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	product1 := &Product1{}
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := &Product2{}
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}

func TestProductFactory_Create(t *testing.T) {
	pf := ProductFactory{}
	product1 := pf.Create(p1)
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := pf.Create(p2)
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}
