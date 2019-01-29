package main

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/scheduler"
	"crawl/LearnGo-crawl/parse/zhengai"
	"crawl/LearnGo-crawl/crawl_distribute/client"
)

func main(){

	itemsave,err:= client.ItemSave(":1234")

	if err!=nil{
		panic(err)
	}
	e:= engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
		itemsave,
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		Parse:engine.NewFuncparse(zhengai.ParseCity,"Parsecity") ,
	})
}


