package main

import (
	"simqo.com/mygospace/learngo/example/crawl/engine"
	"simqo.com/mygospace/learngo/example/crawl/parse"
	"simqo.com/mygospace/learngo/example/crawl/persist"
	"simqo.com/mygospace/learngo/example/crawl/scheduler"
)

func main() {
	e:= engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
		persist.ItemSave(),
	}
	e.Run(engine.Requset{
		Url:"https://book.douban.com",
		ParseFunc:parse.ParseContent,
	})
}

