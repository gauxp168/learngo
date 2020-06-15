package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 均等切割
func equalSplit(num, N int) []int {
	arr := []int{}
	if num%N == 0 {
		for i := 0; i < N; i++ {
			arr = append(arr, num/N)
		}
	}else {
		evg := (num-num%N)/(N-1)
		for i := 0; i < N-1; i++ {
			arr = append(arr, evg)
			num-=evg
		}
		arr = append(arr, num)
	}
	return arr
}

func main() {
	fmt.Println(equalSplit(108, 10))

	// 数据切割
	// 1. 读取需要切割的数据
	sourcefile := ""
	file, err := os.Open(sourcefile)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	lines := getLines(sourcefile)
	arr := equalSplit(lines, 7)
	for i := 0; i < len(arr); i++ {
		// 创建新文件
		newfile := ""+"data_"+ strconv.Itoa(i) +".txt"
		create, err := os.Create(newfile)
		if err != nil {
			fmt.Errorf("error:%v", err)
		}
		defer create.Close()
		writer := bufio.NewWriter(create)
		for j := 0; j < arr[i]; j++ {
			line, _, _ := reader.ReadLine()
			fmt.Fprintln(writer, line)
		}
	}
}

func getLines(sourcefile string) int {
	file, err := os.Open(sourcefile)
	if err != nil {
		fmt.Errorf("error:%v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	sum :=0
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("error:%v", err)
		}
		sum++
	}
	return sum
}
