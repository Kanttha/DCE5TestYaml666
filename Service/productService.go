package Service

import "GoTest1/Model"

type ProductService struct {
	products   []Model.Product
	productNum int //客户数量

}

func (p *ProductService) Create(product Model.Product) bool {
	for _, m := range p.products {
		if product.GetName() == m.GetName() {
			return false
		}
	}
	p.productNum += 1
	product.SetId(p.productNum)
	p.products = append(p.products, product)
	return true
}
func (p *ProductService) GetProducts() []Model.Product {
	return p.products
}

func (p *ProductService) SetName(number int, i string) bool {
	for k := 0; k < len(p.products); k++ {
		if p.products[k].GetName() == i {
			return false
		}
	}
	for j := 0; j < len(p.products); j++ {
		if number == p.products[j].GetId() {
			p.products[j].SetName(i)
			return true
		}
	}
	return false
}
func (p *ProductService) SetPrice(number int, i int64) bool {
	for j := 0; j < len(p.products); j++ {
		if number == p.products[j].GetId() {
			p.products[j].SetPrice(i)
			return true
		}
	}
	return false
}
func (p *ProductService) SetQuantity(number int, i int) bool {
	for j := 0; j < len(p.products); j++ {
		if number == p.products[j].GetId() {
			p.products[j].SetQuantity(i)
			return true
		}
	}
	return false
}
func (p *ProductService) SetKind(number int, i string) bool {
	for j := 0; j < len(p.products); j++ {
		if number == p.products[j].GetId() {
			p.products[j].SetKind(i)
			return true
		}
	}
	return false
}

func (p *ProductService) Delete(number int) bool {
	for i := 0; i < len(p.products); i++ {
		if number == p.products[i].GetId() {
			p.products = append(p.products[:i], p.products[i+1:]...)
			return true
		}
	}
	return false

}
func NewProductService() *ProductService {
	productService := &ProductService{}
	productService.productNum = 1
	product := Model.NewProduct(1, "水泥", 6666, 1, "建材")
	productService.products = append(productService.products, product)
	return productService
}
