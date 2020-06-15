package main

import (
	"math/rand"
	"time"
)

func createArr(length int) []int {
	var list []int
	// 以时间戳为种子生成随机数，保证每次运行数据不重复
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 设置种子，不然每次都会随机成0
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		list = append(list, r.Intn(1000))
	}
	return list
}

func BinInsertSort(list []int, start int, end int, current int) int {
	if start >= end {
		if list[start] < list[current] {
			return current
		}else {
			return start
		}
	}
	center := (end-start)/2 + start
	if list[center] > list[current] {
		return BinInsertSort(list, start, center, current)
	}else {
		return BinInsertSort(list, center+1, end, current)
	}
}

func Sort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		p := BinInsertSort(arr, 0, i-1, i)
		if p != i {
			temp := arr[i]
			for j := i; j > p ; j-- {
				arr[j],arr[j-1] = arr[j-1],arr[j]
			}
			arr[p] = temp
		}
	}
	return arr
}

func main() {
	
}
