package main

import (
	"math/rand"
	"sync"
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

func Merge(left, right []int) []int {
	var data []int
	i,j := 0,0
	l,r := len(left), len(right)
	for i < l && j < r {
		if left[i] < right[j] {
			data = append(data, left[i])
			i++
		}else if left[i] > right[j] {
			data = append(data, right[j])
			j++
		}else {
			data = append(data, left[i])
			i++
			data = append(data, right[j])
			j++
		}
	}
	for i < l {
		data = append(data, left[i])
		i++
	}
	for j < r {
		data = append(data, right[j])
		j++
	}
	return data
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	i:= len(arr)/2
	var wg sync.WaitGroup
	var left, right []int
	wg.Add(2)
	go func() {
		left = MergeSort(arr[0:i])
		wg.Done()
	}()
	go func() {
		right = MergeSort(arr[i:])
		wg.Done()
	}()
	wg.Wait()
	ret := Merge(left, right)
	return ret
}

func main() {
	
}
