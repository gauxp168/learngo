package main

import "fmt"

//goto

func main() {


	//跳出多层循环
	for i:=0; i<10; i++ {
		for j :='A'; j<'Z'; j++ {
			if j == 'C' {
				goto  xx
			}
			fmt.Println(j)
		}
	}
	xx:// label
		fmt.Println("over")
}
