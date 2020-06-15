package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map

func main() {
	// 1. 定义map
	var m1 map[string]int
	// 还没有初始化（没有在内存中开辟空间）
	fmt.Println(m1==nil)

	// 2. 初始化map
	// 要估算好该map容量，避免在程序运行期间再动态扩容
	m1 = make(map[string]int, 10)
	// 3. 赋值
	m1["理想"] = 22
	m1["理性"] = 33
	fmt.Println(m1)
	fmt.Println(m1["理想"])

	// 4. map查找key
	// 如果不存在这个key拿到对应值类型的零值
	fmt.Println(m1["理财"])
	// 约定成俗用ok接收返回的布尔值
	val,ok := m1["理财"]
	if ok {
		fmt.Println(val)
	}else {
		fmt.Println("此key不存在")
	}

	// 5. map 遍历
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	// 6. map 删除
	delete(m1, "理想")
	fmt.Println(m1)
	delete(m1,"理财")

	// 7. map存放为无序
	// 对map进行有序输出  key的排序
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i:= 0; i<100; i++ {
		//生成stu开头的字符串
		key := fmt.Sprintf("stu%02d", i)
		//生成0~99的随机整数
		val := rand.Intn(100)
		scoreMap[key] = val
	}
	fmt.Println(scoreMap)
	//取出map中的所有key存入切片keys
	keys := make([]string,0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	// 8. map&slice组合
	// 元素类型为map的切片
	var sm []map[int]string
	sm = make([]map[int]string, 1, 100)
	// 需要对内部的map初始化
	// 并且每一个slice元素内部的map都需要单独初始化
	sm[0] = make(map[int]string,1)
	sm[0][10] = "hello"
	fmt.Println(sm)

	// 值为切片类型的map
	var ms = make(map[string][]int, 10)
	ms["test"] = []int{1,2,3}
	fmt.Println(ms)
}
