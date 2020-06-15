package main

import "fmt"

func InsertSort(arr []int) []int {
	if len(arr) <=1 {
		return arr
	}else {
		for i := 1; i < len(arr); i++ {
			backup := arr[i]
			j := i-1
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j] = backup
		}
		return arr
	}
}

func main() {
	arr := []int{1,13,24,5,2,7,15,9,43,80,12}
	fmt.Println(InsertSort(arr))
}
