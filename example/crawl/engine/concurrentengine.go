package engine

import (
	"fmt"
	"log"
	"simqo.com/mygospace/learngo/example/crawl/fetcher"
)

type Scheduler interface {
	Submit( Requset )
	Run()
	WorkReady(chan Requset)
	WorkChan() (chan Requset)
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItenChan chan interface{}
}

func (c ConcurrentEngine) Run(seeds ...Requset)  {
	out := make(chan ParseResult)
	c.Scheduler.Run()
	for i:=0; i<c.WorkCount; i++ {
		CreateWork(c.Scheduler.WorkChan(), out, c.Scheduler)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	for  {
		result := <-out
		for _, item := range result.Items {
			go func() {
				c.ItenChan<-item
			}()
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}
func CreateWork(in chan Requset, out chan ParseResult, s Scheduler) {
	go func() {
		for  {
			s.WorkReady(in)
			request := <-in
			result,err := worker(request)
			if err != nil {
				continue
			}
			out<- result
		}
	}()
}
func worker(r Requset) (ParseResult,error) {
	fmt.Printf("fetch url:%s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch Error:%s", r.Url)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}