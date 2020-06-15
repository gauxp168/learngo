package main

type Operator interface {
	SetA(int)
	SetB(int)
	Result()int
}

type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a,b int
}

func (o *OperatorBase) SetA(a int)  {
	o.a = a
}

func (o *OperatorBase) SetB(b int)  {
	o.b = b
}

/*func (o *OperatorBase) Result() int {
	return o.a +o.b
}*/

type PlusOperatorFactory struct {

}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase:&OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

type MinusOperatorFactory struct {

}

type MinusOPerator struct {
	*OperatorBase
}

func (o MinusOPerator) Result() int {
	return o.a - o.b
}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOPerator{
		OperatorBase:&OperatorBase{},
	}
}

func main() {
	
}
