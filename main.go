package main

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/scheduler"
	"crawl/LearnGo-crawl/parse/zhengai"
)

func main(){
	e:= engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:zhengai.ParseCity,
	})
}




