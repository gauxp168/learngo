package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func HandleError(err error, why string)  {
	if err != nil {
		fmt.Println(why, err)
	}
}

func DownloadFile(url string, filename string) (ok bool)  {
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	filename = "E:/Go/src/simqo.com/mygospace/learngo/example/crawler/goroutine" + filename
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return  false
	}else {
		return true
	}
}

/*
// 并发爬思路：
// 1.初始化数据管道
// 2.爬虫写出：26个协程向管道中添加图片链接
// 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// 4.下载协程：从管道里读取链接并下载
*/

var (
	chanImageUrls chan string
	wg sync.WaitGroup
	chanTask chan string
	reImg = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func GetPageStr(url string) (pagestr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll ")
	pagestr = string(bytes)
	return
}

func GetImgs(url string) (urls []string) {
	pagestr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pagestr, -1)
	for _, result := range results {
		url = result[0]
		urls = append(urls, url)
	}
	return
}

func GetImgUrls(url string)  {
	urls := GetImgs(url)
	for _, url := range urls {
		chanImageUrls <- url
	}
	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	wg.Done()
}

func CheckOK()  {
	var count  int
	for  {
		url := <- chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	wg.Done()
}

func GetFilenameFromUrl(url string) (filename string) {
	lastIndex := strings.LastIndex(url,"/")
	filename = url[lastIndex+1:]
	timePerfix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePerfix+"_"+filename
	return
}

func DownloadImg()  {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	wg.Done()
}

func main()  {
	chanImageUrls = make(chan string, 100000)
	chanTask = make(chan string, 26)
	for i:=0; i<27; i++ {
		wg.Add(1)
		go GetImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	wg.Add(1)
	go CheckOK()
	for i:=0; i<5; i++ {
		wg.Add(1)
		go DownloadImg()
	}
	wg.Wait()
}