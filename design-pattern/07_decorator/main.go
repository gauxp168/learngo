package main

import "fmt"

type Component interface {
	Calc() int
}

type ConcreteComponent struct {

}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func (m *MulDecorator) Calc() int {
	return m.Component.Calc() * m.num
}

func NewMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component:c,
		num:num,
	}
}

type AddDecorator struct {
	Component
	num int
}

func (a *AddDecorator) Calc() int {
	return a.Component.Calc() + a.num
}

func NewAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component:c,
		num:num,
	}
}

func main() {
	var c Component = &ConcreteComponent{}
	c = NewAddDecorator(c, 10)
	calc := c.Calc()
	fmt.Printf("ret %d\n", calc)
}
