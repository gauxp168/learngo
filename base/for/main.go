package main

import "fmt"

//for 循环

func main() {
	// 1
	//for i := 0 ; i < 10 ; i++  {
	//	fmt.Println(i)
	//}
	//
	//// 2
	//var  i  = 1
	//for ;i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//// 3
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}
	// 4
	//
	//for  {
	//	fmt.Println(123)
	//}

	// 5
	s := "hello"
	for k, v := range s {
		fmt.Println(v)
		fmt.Printf("%d %c\n", k, v)
	}

	// 6
	for i:=0;i< 10 ; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// 7
	for i:= 1; i<10; i++ {
		if i==5 {
			continue
		}
		fmt.Println(i)
	}
}
