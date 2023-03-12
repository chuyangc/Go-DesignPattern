package DecoratorPattern

import "testing"

func TestComponent1_Operate(t *testing.T) {
	c1 := Component1{}
	c1.Operate()
}

func TestDecorator1_Operate(t *testing.T) {
	d1 := Decorator1{}
	c1 := Component1{}
	d1.c = &c1
	d1.Operate()
}
