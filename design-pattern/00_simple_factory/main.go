package main

import "fmt"

type API interface {
	Say(name string) string
}

type hiAPI struct {

}

func (hi hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s\n", name)
}

type helloAPI struct {

}

func (he helloAPI) Say(name string) string {
	return fmt.Sprintf("hello, %s\n", name)
}

// 简单工厂
func NewAPI(str string) API {
	if str == "hi" {
		return hiAPI{}
	}else if str == "hello" {
		return helloAPI{}
	}
	return nil
}

func main() {
	
}
