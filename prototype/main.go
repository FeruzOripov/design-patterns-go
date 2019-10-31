package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/copier"
)

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	primitive         interface{}
	component         interface{}
	circularReferance ComponentWithBackReference
}

func (p ConcretePrototype) GetPrimitive() interface{} {
	return p.primitive
}

func (p *ConcretePrototype) SetPrimitive(value interface{}) {
	p.primitive = value
}

func (p ConcretePrototype) GetComponent() interface{} {
	return p.component
}

func (p *ConcretePrototype) SetComponent(value interface{}) {
	p.component = value
}

func (p ConcretePrototype) GetCircularReference() ComponentWithBackReference {
	return p.circularReferance
}

func (p *ConcretePrototype) SetCircularReference(value ComponentWithBackReference) {
	p.circularReferance = value
}

func (p *ConcretePrototype) Clone() Prototype {
	copier.Copy(p.component, p.component)
	copier.Copy(p.circularReferance, p.circularReferance)
	p.circularReferance.prototype = p
	return copier.Copy(p, p)
}

type ComponentWithBackReference struct {
	prototype Prototype
}

func (c ComponentWithBackReference) GetPrototype() Prototype {
	return c.prototype
}

func (c *ComponentWithBackReference) SetPrototype(value Prototype) {
	c.prototype = value
}

func main() {
	p1 := ConcretePrototype{}
	p1.primitive = 245
	p1.component = time.Now()
	p1.circularReferance = ComponentWithBackReference{prototype: p1}

	p2 := p1.Clone()

	if p1.primitive == p2.primitive {
		fmt.Println("Primitive field values have been carried over to a clone. Yay!")
	}
}
