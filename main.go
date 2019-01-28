package main

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/parse"
	"crawl/LearnGo-crawl/scheduler"
)

func main(){
	e:= engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
	}

	e.Run(engine.Request{
		Url:"https://book.douban.com",
		ParseFunc:parse.ParseTag,
	})
}




