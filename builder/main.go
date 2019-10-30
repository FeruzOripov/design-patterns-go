package main

import (
	"fmt"
	"strings"
)

type Builder interface {
	Product() Product1
	producePartA()
	producePartB()
	producePartC()
}

type ConcreteBuilder1 struct {
	product Product1
}

func (c ConcreteBuilder1) Reset() {
	c.product = Product1{}
}

func (c ConcreteBuilder1) Product() Product1 {
	product := c.product
	c.Reset()
	return product
}

func (c ConcreteBuilder1) producePartA() {
	c.product.Add("PartA1")
}

func (c ConcreteBuilder1) producePartB() {
	c.product.Add("PartB1")
}

func (c ConcreteBuilder1) producePartC() {
	c.product.Add("PartC1")
}

type Product1 struct {
	parts []string
}

func (p *Product1) Add(part string) {
	p.parts = append(p.parts, part)
}

func (p *Product1) ListParts() {
	fmt.Printf("Product parts: {%s}", strings.Join(p.parts, ", "))
}

type Director struct {
	builder Builder
}

func (d Director) GetBuilder() Builder {
	return d.builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) BuildMinimalViableProduct() {
	d.builder.producePartA()
}

func (d *Director) BuildFullFeaturedProduct() {
	d.builder.producePartA()
	d.builder.producePartB()
	d.builder.producePartC()
}

func main() {
	var director Director
	var builder ConcreteBuilder1
	director.builder = &builder

	fmt.Println("Standart basic product: ")
	director.BuildMinimalViableProduct()
	builder.product.ListParts()

	fmt.Println()

	fmt.Println("Standart full featured product: ")
	director.BuildFullFeaturedProduct()
	builder.product.ListParts()

	fmt.Println()

	fmt.Println("Custom product: ")
	builder.producePartA()
	builder.producePartB()
	builder.product.ListParts()
}
