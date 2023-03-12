package FactoryMethod

import (
	"fmt"
	"testing"
)

func TestProduct1Factory_Create(t *testing.T) {
	p1f := Product1Factory{}
	product1 := p1f.Create()
	product1.SetName("product1")
	fmt.Println(product1.GetName())
}

func TestProduct2Factory_Create(t *testing.T) {
	p2f := Product2Factory{}
	product2 := p2f.Create()
	product2.SetName("product2")
	fmt.Println(product2.GetName())
}
