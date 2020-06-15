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

func QuickGoSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	var wg sync.WaitGroup
	c:=data[0]
	var left, mid, right []int
	mid = append(mid, c)
	for k, v := range data {
		if k == 0 {
			continue
		}
		if c > v {
			left = append(left,v)
		}else if c < v {
			right = append(right, v)
		}else {
			mid = append(mid, v)
		}
	}
	wg.Add(2)
	go func() {
		left= QuickGoSort(left)
		wg.Done()
	}()

	go func() {
		right = QuickGoSort(right)
		wg.Done()
	}()
	wg.Wait()
	var ret []int
	if len(left) > 0 {
		ret = append(ret, left...)
	}
	ret = append(ret, mid...)
	if len(right) > 0 {
		ret = append(ret, right...)
	}
	return ret
}

func main() {
	
}
