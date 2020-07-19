package demo1

import (
	"fmt"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	product1 := &Product1{}
	product1.SetName("jerry")
	fmt.Println(product1.GetName())

	product2 := &Product2{}
	product2.SetName("tom")
	fmt.Println(product2.GetName())
}

func TestProductFactory_Create(t *testing.T) {
	productFactory := productFactory{}
	product1 := productFactory.Create(p1)
	product1.SetName("bob")
	fmt.Println(product1.GetName())

	product2 := productFactory.Create(1)
	product2.SetName("zhang")
	fmt.Println(product2.GetName())
}
