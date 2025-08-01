package main

import (
	"GoTest1/Model"
	"GoTest1/Service"
	"fmt"
)

type productView struct {
	key            string //接收输入
	loop           bool   //是否循环
	productService *Service.ProductService
	product        Model.Product
}

// 显示货物信息
func (p *productView) list() {
	products := p.productService.GetProducts()
	fmt.Println("-----------货 物 列 表-----------")
	fmt.Println("编号\t\t名称\t\t价格\t\t数量\t\t种类")
	for _, product := range products {
		fmt.Println(product.ToString())
	}
}

// 添加货物信息
func (p *productView) add() {
	fmt.Println("请输入货物名称,货物价格，货物数量，货物种类：")
	var (
		name     string
		price    int64
		quantity int
		kind     string
	)
	fmt.Print("请输入 名称 价格 数量 种类（空格分隔）：")
	if _, err := fmt.Scanln(&name, &price, &quantity, &kind); err != nil {
		fmt.Println("输入有误：", err)
		return
	}
	p.product.SetName(name)
	p.product.SetPrice(price)
	p.product.SetQuantity(quantity)
	p.product.SetKind(kind)
	if !p.productService.Create(p.product) {
		fmt.Println("仓库中已经存在同名货物！")
		return
	}
	fmt.Println("成功添加货物信息！")
	p.list()
}

// 修改货物信息
func (p *productView) update() {
	fmt.Println("请输入要修改的货物编号:")
	var number int
	if _, err := fmt.Scanln(&number); err != nil {
		fmt.Println("输入有误：", err)
		return
	}
	fmt.Println("请输入需要修改的内容对应的数字(1:货物名称,2:货物价格，3:货物数量，4:货物种类)：")
	var string1 int
	if _, err := fmt.Scanln(&string1); err != nil && (string1 < 1 || string1 > 4) {
		fmt.Println("输入有误：", err)
		return
	}
	fmt.Println("要修改为:")
	switch string1 {
	case 1:
		var i string
		if _, err := fmt.Scanln(&i); err != nil {
			fmt.Println("输入有误：", err)
			return
		}
		if !p.productService.SetName(number, i) {
			fmt.Println("存在同名货物或操作错误，失败！")
			return
		}
		fmt.Println("成功修改货物信息！")
	case 2:
		var i int64
		if _, err := fmt.Scanln(&i); err != nil {
			fmt.Println("输入有误：", err)
			return
		}
		if !p.productService.SetPrice(number, i) {
			fmt.Println("操作错误，失败！")
			return
		}
		fmt.Println("成功修改货物信息！")
	case 3:
		var i int
		if _, err := fmt.Scanln(&i); err != nil {
			fmt.Println("输入有误：", err)
			return
		}
		if !p.productService.SetQuantity(number, i) {
			fmt.Println("操作错误，失败！")
			return
		}
		fmt.Println("成功修改货物信息！")

	case 4:
		var i string
		if _, err := fmt.Scanln(&i); err != nil {
			fmt.Println("输入有误：", err)
			return
		}
		if !p.productService.SetKind(number, i) {
			fmt.Println("操作错误，失败！")
			return
		}
		fmt.Println("成功修改货物信息！")
	default:
		fmt.Println("请重新输入！")
	}

	p.list()
}

// 删除货物
func (p *productView) remove() {
	fmt.Println("请输入要删除的货物编号:")
	var number int
	if _, err := fmt.Scanln(&number); err != nil {
		fmt.Println("输入有误：", err)
		return
	}
	if !p.productService.Delete(number) {
		fmt.Println("输入有误：")
		return
	}
	fmt.Println("成功删除货物信息！")
	p.list()
}

// 主界面
func (p *productView) mainMenu() {
	for {
		fmt.Println("-----------仓 库 管 理 系 统-----------")
		fmt.Println("			1 添 加 货 物")
		fmt.Println("			2 修 改 货 物")
		fmt.Println("			3 删 除 货 物")
		fmt.Println("			4 货 物 总 览")
		fmt.Println("			5 退	  出")
		fmt.Println("请选择1-5：")
		_, err := fmt.Scanln(&p.key)
		if err != nil {
			fmt.Println(err)
		}
		switch p.key {
		case "1":
			p.add()
		case "2":
			p.update()
		case "3":
			p.remove()
		case "4":
			p.list()
		case "5":
			p.loop = false
		default:
			fmt.Println("请重新输入！")
		}
		if !p.loop {
			break
		}
	}
	fmt.Println("退出")
}

func main() {
	productView := productView{
		key:  "",
		loop: true,
	}
	productView.productService = Service.NewProductService()
	//主页
	productView.mainMenu()
}
