package main

import (
	"container/list"
	"math/rand"
	"time"
)

// 归并算法

// 归并排序的简单归并
func Merge(arr1 []int, arr2 []int) []int {
	allarr := []int{}
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			allarr = append(allarr, arr1[i])
			i++
		}else if arr1[i] > arr2[j] {
			allarr = append(allarr, arr1[j])
			j++
		}else {
			allarr = append(allarr, arr1[i])
			i++
			allarr = append(allarr, arr1[j])
			j++
		}
	}
	for i < len(arr1) {
		allarr = append(allarr, arr1[i])
		i++
	}
	for j < len(arr2) {
		allarr = append(allarr, arr1[j])
		j++
	}
	return allarr
}

// 归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	i:= len(arr)/2
	left := MergeSort(arr[0:i])
	right := MergeSort(arr[i:])
	ret := Merge(left, right)
	return ret
}

// 基于栈  先入后出
func listMerge(arr []string) string {
	mylist := list.New()
	for i := 0; i < len(arr); i++ {
		mylist.PushBack(arr[i])
	}
	for mylist.Len() !=1 {
		e1 := mylist.Back()
		mylist.Remove(e1)
		e2 := mylist.Back()
		mylist.Remove(e2)

		if e1 != nil && e2 != nil {
			v1 ,_ := e1.Value.(string)
			 v2,_ := e2.Value.(string)
			 v3:= v1+v2
			 mylist.PushBack(v3)
		}else if e1 != nil && e2 == nil {
			v1,_:=e1.Value.(string)
			mylist.PushBack(v1)
		}else if e1 == nil && e2 == nil {
			break
		}else {
			break
		}
	}
	return mylist.Back().Value.(string)
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

func main() {
	
}

