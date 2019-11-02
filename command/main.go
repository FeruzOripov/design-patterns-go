package main

import "fmt"

type Command interface {
	Execute()
}

type SimpleCommand struct {
	Payload string
}

func (s SimpleCommand) Execute() {
	fmt.Printf("SimpleCommand: See, I can do simple things like printing (%s)\n", s.Payload)
}

type ComplexCommand struct {
	receiver Receiver
	a        string
	b        string
}

func (c ComplexCommand) Execute() {
	fmt.Println("ComplexCommand: Complex stuff should be done by a receiver object")
	c.receiver.doSomething(c.a)
	c.receiver.doSomethingElse(c.b)
}

type Receiver struct{}

func (r Receiver) doSomething(a string) {
	fmt.Printf("Receiver: Working on (%s)\n", a)
}

func (r Receiver) doSomethingElse(b string) {
	fmt.Printf("Receiver: Also working on (%s)\n", b)
}

type Invoker struct {
	onStart  Command
	onFinish Command
}

func (i *Invoker) setOnStart(command Command) {
	i.onStart = command
}

func (i *Invoker) setOnFinish(command Command) {
	i.onFinish = command
}

func (i *Invoker) doSomethingImportant() {
	fmt.Println("Invoker: Does anybody want something done before I begin ?")
	i.onStart.Execute()

	fmt.Println("Invoker: ...doing something really important...")

	fmt.Println("Invoker: Does anybody want something done after I finish ?")
	i.onFinish.Execute()
}

func main() {
	invoker := Invoker{}
	invoker.setOnStart(SimpleCommand{"Say Hi!"})
	receiver := Receiver{}
	invoker.setOnFinish(ComplexCommand{receiver, "Send email", "Save report"})

	invoker.doSomethingImportant()
}
