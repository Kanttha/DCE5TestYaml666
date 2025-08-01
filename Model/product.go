package Model

import "fmt"

type Product struct {
	id       int
	name     string //货物名称
	price    int64  //货物价格
	quantity int    //货物数量
	kind     string //货物种类
}

func GetProduct() Product {
	return Product{}
}

func NewProduct(Id int, name string, price int64, quantity int, kind string) Product {
	var p Product
	p.id = Id
	p.name = name
	p.price = price
	p.quantity = quantity
	p.kind = kind
	return p
}
func (p *Product) GetId() int {
	return p.id
}
func (p *Product) GetName() string {
	return p.name
}
func (p *Product) GetPrice() int64 {
	return p.price
}
func (p *Product) GetQuantity() int {
	return p.quantity
}
func (p *Product) GetKind() string {
	return p.kind
}
func (p *Product) SetId(id int) {
	p.id = id
}
func (p *Product) SetName(name string) {
	p.name = name
}
func (p *Product) SetPrice(price int64) {
	p.price = price
}
func (p *Product) SetQuantity(quantity int) {
	p.quantity = quantity
}
func (p *Product) SetKind(kind string) {
	p.kind = kind
}
func (p *Product) ToString() string {
	return fmt.Sprintf("%v\t\t%v\t\t%v\t\t%v\t\t%v", p.id, p.name, p.price, p.quantity, p.kind)
}
