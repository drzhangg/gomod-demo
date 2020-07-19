package demo1

type Product interface {
	SetName(name string)
	GetName() string
}

type Product1 struct {
	name string
}

func (p *Product1) SetName(name string) {
	p.name = name
}

func (p *Product1) GetName() string {
	return "产品1的name为：" + p.name
}

type Product2 struct {
	name string
}

func (p *Product2) SetName(name string) {
	p.name = name
}

func (p *Product2) GetName() string {
	return "产品2的name为：" + p.name
}

type productType int

const (
	p1 productType = iota
	p2
)

//实现简单工厂类

type productFactory struct {
}

func (pf productFactory) Create(productType productType) Product {
	if productType == 1 {
		return &Product1{}
	}

	if productType == 2 {
		return &Product2{}
	}
	return nil
}
