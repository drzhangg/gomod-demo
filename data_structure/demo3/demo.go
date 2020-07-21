package demo3

type Product interface {
	SetName(name string)
	GetName() string
}

type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return p1.name
}

type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return p2.name
}

//实现一个抽象工厂
type ProductFactory interface {
	Create() Product
}

type ProductFactory1 struct {
}

func (p1 *ProductFactory1) Create() Product {
	return &Product1{}
}

type ProductFactory2 struct {
}

func (p2 *ProductFactory2) Create() Product {
	return &Product2{}
}
