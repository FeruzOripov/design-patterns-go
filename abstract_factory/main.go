package main

import "fmt"

type AbstractFactory interface {
	createProductA() AbstractProductA
	createProductB() AbstractProductB
}

type ConcreteFactory1 struct{}

func (c ConcreteFactory1) createProductA() AbstractProductA {
	return ConcreteProductA1{}
}

func (c ConcreteFactory1) createProductB() AbstractProductB {
	return ConcreteProductB1{}
}

type ConcreteFactory2 struct{}

func (c ConcreteFactory2) createProductA() AbstractProductA {
	return ConcreteProductA2{}
}

func (c ConcreteFactory2) createProductB() AbstractProductB {
	return ConcreteProductB2{}
}

type AbstractProductA interface {
	usefulFunctionA() string
}

type ConcreteProductA1 struct{}

func (c ConcreteProductA1) usefulFunctionA() string {
	return "The result of the product A1."
}

type ConcreteProductA2 struct{}

func (c ConcreteProductA2) usefulFunctionA() string {
	return "The result of the product A2."
}

type AbstractProductB interface {
	usefulFunctionB() string
	anotherUsefulFunctionB(AbstractProductA) string
}

type ConcreteProductB1 struct{}

func (C ConcreteProductB1) usefulFunctionB() string {
	return "The result of the product B1."
}

func (c ConcreteProductB1) anotherUsefulFunctionB(collaborator AbstractProductA) string {
	result := collaborator.usefulFunctionA()
	return fmt.Sprintf("The result of the B1 collaborating with the %s", result)
}

type ConcreteProductB2 struct{}

func (c ConcreteProductB2) usefulFunctionB() string {
	return "The result of the product B2"
}

func (c ConcreteProductB2) anotherUsefulFunctionB(collaborator AbstractProductA) string {
	result := collaborator.usefulFunctionA()
	return fmt.Sprintf("The result of the B2 collaborating with the %s", result)
}

func clientCode(factory AbstractFactory) {
	productA := factory.createProductA()
	productB := factory.createProductB()

	fmt.Println(productB.usefulFunctionB())
	fmt.Println(productB.anotherUsefulFunctionB(productA))
}

func main() {
	fmt.Println("Client: Testing client code with the first factory type:")
	clientCode(ConcreteFactory1{})

	fmt.Println()

	fmt.Println("Client: Testing the same client code with the second factory type:")
	clientCode(ConcreteFactory2{})
}
