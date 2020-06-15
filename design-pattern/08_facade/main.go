package main

import "fmt"

type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct {

}

func (*aModuleImpl) TestA() string {
	return "a module running"
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleimpl struct {

}

func (*bModuleimpl) TestB() string {
	return "b module running"
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleimpl{}
}

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aret := a.a.TestA()
	bret := a.b.TestB()
	return fmt.Sprintf("%s\n%s\n", aret, bret)
}

func NewAPI() API {
	return &apiImpl{
		a:NewAModuleAPI(),
		b:NewBModuleAPI(),
	}
}

func main() {
	api := NewAPI()
	test := api.Test()
	fmt.Printf(test)
}
