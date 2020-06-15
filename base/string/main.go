package main

import (
	"fmt"
	"strings"
)
// strings ：https://studygolang.com/pkgdoc   or    https://www.godoc.org/strings
// string
func main() {
	// \ 本来是具有特殊含义的，我应该告诉程序我写的\就是一个单纯的\
	path := "D:\\Go\\src\\code.oldboyedu.com\\studygo\\day01"
	fmt.Println(path)

	s := "i am ok!"
	fmt.Println(s)

	s2 := `
世情薄
				人情恶
		雨送黄昏花易落
`

	fmt.Println(s2)

	s3 := `D:\Go\src\code.oldboyedu.com\studygo\day01`
	fmt.Println(s3)

	// 字符串相关操作
	// 1. len()
	fmt.Println(len(s3))

	// 2. 字符串拼接
	name := "理性"
	world := "我们的时间"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	// 3. 分隔
	ret := strings.Split(s3,"\\")
	fmt.Println(ret)

	// 4. 包含
	fmt.Println(strings.Contains(ss,"理性"))
	fmt.Println(strings.Contains(ss,"理想"))

	// 5. 前缀
	fmt.Println(strings.HasPrefix(ss,"理想"))

	// 6. 后缀
	fmt.Println(strings.HasSuffix(ss,"理想"))

	// 7. 位置index
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	//8. 拼接 (针对切片slice)
	fmt.Println(strings.Join(ret, "+"))

}
