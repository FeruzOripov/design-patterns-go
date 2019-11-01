package main

import "fmt"

type Subject interface {
	request()
}

type RealSubject struct{}

func (r *RealSubject) request() {
	fmt.Println("RealSubject: Handling request.")
}

type Proxy struct {
	realSubject Subject
}

func (p *Proxy) request() {
	if p.checkAccess() {
		p.realSubject.request()
		p.logAccess()
	}
}

func (p *Proxy) checkAccess() bool {
	fmt.Println("Proxy: Checking access prior to firing a real request.")
	return true
}

func (p *Proxy) logAccess() {
	fmt.Println("Proxy: Logging the time of request.")
}

func clientCode(subject Subject) {
	subject.request()
}

func main() {
	fmt.Println("Client: Executing the client code with a real subject:")
	var real_subject Subject
	real_subject = &RealSubject{}
	clientCode(real_subject)

	fmt.Println()

	fmt.Println("Client: Executing the same client code with a proxy:")
	var proxy Subject
	proxy = &Proxy{real_subject}
	clientCode(proxy)
}
