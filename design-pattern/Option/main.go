package main

import "fmt"

// 引入问题: 现有一个结构体，再声明一个用于初始化该结构体的构造方法， 进行对结构体字段赋值

/*
选项设计模式实现
声明一个 Option 变量类型，类型为一个函数类型，函数类型接收一个指针类型的 Options 结构体
构造函接受一个可变长度的 Option 变量类型
之后没增加一个字段创建一个指定格式类型的函数，传入构造函数即可初始化结构体
*/

type Options struct {
	str1 string
	str2 string
	int1 int
	int2 int
}

type Option func(opt *Options)

func NewOptions(opts ...Option) *Options {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func WithOptionStr1(str string) Option {
	return func(opt *Options) {
		opt.str1 = str
	}
}

func WithOptionStr2(str string) Option {
	return func(opt *Options) {
		opt.str2 = str
	}
}

func WithOptionInt1(i int) Option {
	return func(opt *Options) {
		opt.int1 = i
	}
}

func WithOptionInt2(i int) Option {
	return func(opt *Options) {
		opt.int2 = i
	}
}

func main() {
	optionis := NewOptions(
		WithOptionStr2("q"),
		WithOptionStr1("a"),
		WithOptionInt1(1),
		WithOptionInt2(2),
		)
	fmt.Printf("%#v\n", optionis)
}
