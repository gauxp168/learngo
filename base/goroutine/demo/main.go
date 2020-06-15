package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	Id int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func main() {
	// 2个管道
	// 1. job管理管道
	jobChan := make(chan *Job, 128)
	// 2. result管理通道
	resultChan := make(chan *Result, 128)
	// 3. 创建工作池
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum,result.sum)
		}
	}(resultChan)
	var id int
	for {
		id++
		rand.Seed(time.Now().UnixNano())
		r_num := rand.Int()
		job := &Job{
			Id:id,
			RandNum:r_num,
		}
		jobChan <- job
	}


}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan <-chan *Job, resultChan chan<- *Result){
	for i := 0; i<num; i++ {
		go func(jobChan <-chan *Job, resultChan chan<- *Result) {
			for job := range jobChan {
				r_num := job.RandNum
				var sum int
				for r_num != 0{
					tmp := r_num%10
					sum += tmp
					r_num /= 10
				}
				r := &Result{
					job: job,
					sum:sum,
				}
				resultChan<- r
			}
		}(jobChan, resultChan)
	}
}
