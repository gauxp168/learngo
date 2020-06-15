package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
爬虫步骤
明确目标（确定在哪个网站搜索）
爬（爬下内容）
取（筛选想要的）
处理数据（按照你的想法去处理）
*/

// example 1
// 这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作

var(
	reQQEmail = `(\d+)@qq.com`
)

func GetEmail()  {
	// 1. 去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2. 读取页面内容
	pageeBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	pageStr := string(pageeBytes)
	//fmt.Println(pageStr)
	// 3. 过滤数据，过虑邮箱
	compile := regexp.MustCompile(reQQEmail)
	results := compile.FindAllStringSubmatch(pageStr, -1)
	fmt.Println(len(results))
	// 遍历结果
	for _, value := range results {
		fmt.Printf("email: %v\n", value[0])
		fmt.Printf("QQ: %v\n", value[1])
	}
}

func HandleError(err error, why string)  {
	if err != nil {
		fmt.Println(why, err)
	}
}

func main() {
	GetEmail()
}
