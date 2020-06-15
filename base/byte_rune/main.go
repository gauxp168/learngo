package main

import "fmt"

// byte和rune类型
// Go语言中为了处理非ASCII码类型的字符 定义了新的rune类型
func main() {
	s := "hello左右사샤"
	n := len(s)
	fmt.Println(n)

	//for i := 0; i < n ; i++ {
	//	fmt.Println(s[i])
	//	fmt.Printf("%c\n", s[i]) // %c:字符
	//}

	//for _,v := range s {
	//	fmt.Printf("%c\n", v)
	//}

	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转换成了一个rune切片
	// => '白' '萝' '卜'
	fmt.Println(s3)
	s3[0] = '红'
	fmt.Println(string(s3))// 把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红' //rune(int32)
	fmt.Printf("c1 %T , c2 %T\n", c1 ,c2)
	c3 := "H"
	c4 := byte('H') //byte(uint8)
	fmt.Printf("c3 %T, c4 %T\n", c3 ,c4)

	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
