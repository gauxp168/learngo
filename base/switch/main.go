package main

import "fmt"

func main() {
	switch n:=3; n {
	case 1:
		fmt.Println("first")
	case 2:
		fmt.Println("second")
	case 3:
		fmt.Println("thild")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("not number")
	}

	switch i:=4; i {
	case 1,3,5,7,9:
		fmt.Println("奇数")
	case 2,4,6,8:
		fmt.Println("偶数")
	default:
		fmt.Println(i)
	}
}
