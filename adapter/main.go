package main

import "fmt"

type Target interface {
	request() string
}

type Adaptee interface {
	specific_request() string
}

type target struct{}

func (t target) request() string {
	return "Target: The default target's behaviour."
}

type adaptee struct{}

func (a adaptee) specific_request() string {
	return ".eetpadA eht fo roivaheb laicepS"
}

type Adapter interface {
	request() string
}

type adapter struct {
	adapteer Adaptee
}

func (a adapter) request() string {
	var res string
	req := a.adapteer.specific_request()
	for i := len(req) - 1; i >= 0; i-- {
		res += string(req[i])
	}
	return fmt.Sprintf("Adapter: (TRANSLATED) %s", res)
}

func clientCode(target Target) {
	fmt.Println(target.request())
}

func main() {
	fmt.Println("Client: I can work just fine with the Target objects:")
	var target target
	clientCode(target)
	fmt.Println()

	var adaptee adaptee
	fmt.Println("Client: The adaptee class has a weird interface. See, I don't understand it:")
	fmt.Printf("Adaptee: %s", adaptee.specific_request())
	fmt.Printf("\n\n")

	fmt.Println("Client: But I can work with it via the Adapter:")
	var adapter adapter
	adapter.adapteer = adaptee
	clientCode(adapter)
}
