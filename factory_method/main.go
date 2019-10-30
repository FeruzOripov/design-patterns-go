package main

import "fmt"

type Creator interface {
	factory_method() Product
}

func some_operation(c Creator) string {
	// Call the factory method to create a Product object.
	product := c.factory_method()

	// Now, use the product
	result := "Creator: The same creator's code has just worked with " + product.operation()

	return result
}

type ConcreteCreator1 struct{}

func (c ConcreteCreator1) factory_method() Product {
	return ConcreteProduct1{}
}

type ConcreteCreator2 struct{}

func (c ConcreteCreator2) factory_method() Product {
	return ConcreteProduct2{}
}

type Product interface {
	operation() string
}

type ConcreteProduct1 struct{}

func (c ConcreteProduct1) operation() string {
	return "Result of the ConcreteProduct1"
}

type ConcreteProduct2 struct{}

func (c ConcreteProduct2) operation() string {
	return "Result of the ConcreteProduct2"
}

func client_code(creator Creator) {
	fmt.Printf("Client: I'm not aware of the creator's class, but it still works.\n%s", some_operation(creator))
}

func main() {
	fmt.Println("App: Launched with the ConcreteCreator1.")
	client_code(ConcreteCreator1{})
	fmt.Println()

	fmt.Println("App: Launched with the ConcreteCreator2.")
	client_code(ConcreteCreator2{})
	fmt.Println()
}
