package main

import (
	"fmt"
	"sort"
)

// 切片 slice
func main() {
	// 切片定义
	var s1 []int
	var s2 []string
	fmt.Println(s1,s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	// 1
	s1 = []int{1,2,3}
	s2 = []string{"你最","你的","你身"}
	fmt.Println(s1,s2)
	fmt.Println(s1 == nil)
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 2
	ss := make([]int, 0, 10)
	fmt.Printf("len(ss):%d cap(ss):%d\n", len(ss), cap(ss))


	// 有数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := a1[1:5]//// 基于一个数组切割，左包含右不包含，（左闭右开）
	fmt.Println(s3)
	s4 := a1[:4]
	s5 := a1[:]
	s6 := a1[3:]
	fmt.Println(s4, s5, s6)

	// 数组切片的容量是指底层数组的容量
	fmt.Printf("len(s4):%d cap(s4):%d \n", len(s4), cap(s4))
	// 数组切片的容量是底层数组从切片的第一个元素到最后的元素的数量
	fmt.Printf("len(s6):%d cap(s6):%d \n", len(s6), cap(s6))
	// 切片在分割
	s8 := s6[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))

	//切片为引用类型，都指向底层的数组
	fmt.Println(s6)
	a1[6] = 100
	fmt.Println(s6)
	fmt.Println(s8)

	// append 为切片追加元素
	sa1 := []string{"北京", "上哈", "深圳"}
	fmt.Printf("sa1=%v len(sa1)=%d cap(sa1)=%d\n", sa1, len(sa1), cap(sa1))

	// 调用append函数必须用原来的切片变量接收返回值
	// append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	// 必须用变量接收append的返回值
	sa1 = append(sa1, "广州")
	fmt.Printf("sa1=%v len(sa1)=%d cap(sa1)=%d\n", sa1, len(sa1), cap(sa1))
	sa1 = append(sa1, "杭州", "成都")
	fmt.Printf("sa1=%v len(sa1)=%d cap(sa1)=%d\n", sa1, len(sa1), cap(sa1))
	sss := []string{"武汉", "苏州"}
	sa1 = append(sa1, sss...)
	fmt.Printf("sa1=%v len(sa1)=%d cap(sa1)=%d\n", sa1, len(sa1), cap(sa1))

	// copy 切片复制成 一个不基于原来切片底层数组的新切片
	cs1 :=[]int{1,2,3}
	cs2 := cs1
	cs3 := make([]int, 3,3)
	copy(cs3, cs1)
	fmt.Println(cs1,cs2,cs3)
	cs1[0] = 100
	fmt.Println(cs1,cs2,cs3)

	// 切片删除
	// 删除切片中某一位
	cs1 = append(cs1[:1], cs1[2:]...)
	fmt.Println(cs1, cap(cs1))

	// 数组切片中间删除某一位的变化
	x1 := [...]int{1,3,5}
	xx := x1[:]
	fmt.Println(xx, len(xx), cap(xx))
	fmt.Printf("%p\n", &xx[0])
	xx = append(xx[:1], xx[2:]...)
	fmt.Printf("%p\n", &xx[0])
	fmt.Println(x1,xx, len(xx), cap(xx))

	xx[0] = 222
	fmt.Println(x1)

	// 切片排序
	var qq = [...]int{8,4,1,7,2,5,3,9}
	ww := qq[:]
	sort.Ints(ww)
	fmt.Println(ww)

}
