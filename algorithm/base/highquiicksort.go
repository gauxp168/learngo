package main

import (
	"math/rand"
	"time"
)

func BinSearchSort(mylist []int) []int {
	if len(mylist) <= 1 {
		return mylist
	}else {
		for i := 1; i < len(mylist); i++ {
			p := BinInsertSort(mylist, 0, i-1, i)
			if p != i {
				for j := i; j > p ; j-- {
					mylist[j],mylist[j-1] = mylist[j-1],mylist[j]
				}
			}
		}
		return mylist
	}
}

// 对指定数据段排序
func BinSearchSortIndex(mylist []int, start int, end int) []int {
	if end - start<=1 {
		return mylist
	}else {
		for i := start+1; i <= end; i++ {
			p := BinInsertSort(mylist, start, i-1, i)
			if p != i {
				for j := i; j > p ; j-- {
					mylist[j],mylist[j-1] = mylist[j-1],mylist[j]
				}
			}
		}
		return mylist
	}
}
func QuickSortCall(arr []int) []int {
	if len(arr) < 20 {
		return BinSearchSort(arr)
	}else {
		QuickSort(arr, 0, len(arr)-1)
		return arr
	}
}

func Swap(arr []int, i,j int)  {
	arr[i],arr[j] = arr[j],arr[i]
}

func QuickSort(arr []int, left int, right int)  {
	if right - left < 10 {
		BinSearchSortIndex(arr, left, right)
	}else {
		// 快速排序
		// 任何位置，交换到第一位
		Swap(arr, left, rand.Int()%(right-left)+left)
		vdata := arr[left]
		lt := left    // arr[left+1....lt]  < vdata  lt++
		gt := right+1  // arr[gt.....right] > vdata  gt--
		i := left +1   // arr[lt+1....i]  == vdata   i++
		for i < gt {
			if arr[i] < vdata {
				Swap(arr, i, lt+1)
				lt++
				i++
			}else if arr[i] > vdata {
				Swap(arr,i,gt-1)
				gt--
			}else {
				i++
			}
		}
		Swap(arr, left,lt)
		QuickSort(arr, left, lt-1)
		QuickSort(arr, gt, right)
	}
}

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


func main() {

	/*mylist := createArr(10)
	mylist = QuickSortCall(mylist)*/
}
