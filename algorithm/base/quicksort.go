package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}else {
		tmp := arr[0]
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid = append(mid, tmp)
		for i := 1; i < len(arr); i++ {
			if tmp > arr[i] {
				low = append(low,arr[i])
			}else if tmp < arr[i] {
				high = append(high, arr[i])
			}else {
				mid = append(mid, arr[i])
			}
		}
		low = QuickSort(low)
		high = QuickSort(high)
		ret := append(append(low, mid...), high...)
		return ret
	}
}

func HighQuickSort(arr []int) []int {

}

func main() {
	arr := []int{23,12,43,21,45,64,1,5,86,347,234}
	fmt.Println(QuickSort(arr))
}
