package main

import "fmt"

// 数组

// 存放元素的容器
// 必须指定存放的元素的类型和容量（长度）
// 数组的长度是数组类型的一部分
func main() {
	var a1  [3]bool
	var a2  [5]bool

	fmt.Printf("a1:%T a2%T \n", a1,a2)
	// 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false, 整型和浮点型都是0, 字符串：""
	fmt.Println(a1,a2)
	// 1
	a1 = [3]bool{true, true}
	// 2
	a3 := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(a3)
	// 3
	a4 := [5]int{1:2,3:5}
	fmt.Println(a4)

	//数组的遍历
	citys := [...]string{"北京","上海","深圳"}
	// 1
	for i := 0; i<len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 2
	for k,v := range citys{
		fmt.Println(k,v)
	}

	// 多维数组
	var arr [3][2]int
	arr = [3][2]int{
		[2]int{1,2},
		[2]int{3,4},
		[2]int{5,6},
	}
	fmt.Println(arr)

	for _,v1 := range arr{
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组为值类型
	b1 := [3]int{1,2,3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1,b2)
}
