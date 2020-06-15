package main

import (
	"fmt"
	"reflect"
)

// reflect 反射

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
	fmt.Printf("type :%T\n", v)
	fmt.Printf("type name:%v type kind:%v\n", v.Name(), v.Kind())

}

type cat struct {

}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())
	k := v.Kind()
	switch k {
	case reflect.Float32:
		fmt.Printf("this is type float32, value:%f type:%T\n", v.Float(),v.Float())

	case reflect.Float64:
		fmt.Printf("this is type float64, value:%f\n", v.Float())

	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind()== reflect.Int64)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}
func reflectSetValueElem(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

type student struct {
	Name string `json:"name" info:"hello"`
	Score int `json:"score" info:"test"`
}


func main() {
	var a float32 = 3.14
	reflectType(a)
	var b float64 = 2.34
	reflectType(b)
	var c  = struct {

	}{}
	reflectType(c)
	var d = cat{}
	reflectType(d)
	reflectValue(a)

	var f  int64 = 100
	reflectSetValue(&f)
	fmt.Println(f)
	reflectSetValueElem(&f)
	fmt.Println(f)

	stu := student{
		Name:"小河",
		Score:88,
	}
	of := reflect.TypeOf(stu)
	fmt.Println(of.Name(), of.Kind())

	fmt.Println(of.NumField())
	for i := 0; i < of.NumField(); i++ {
		field := of.Field(i)
		fmt.Printf("namee:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
}
