package main

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/scheduler"
	"crawl/LearnGo-crawl/parse/zhengai"
	"crawl/LearnGo-crawl/crawl_distribute/client"
	client2 "crawl/LearnGo-crawl/crawl_distribute/work/client"
)

func main(){

	itemsave,err:= client.ItemSave(":1234")

	process,err:=  client2.CreateProcess()
	if err!=nil{
		panic(err)
	}
	e:= engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
		itemsave,
		process,
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		Parse:engine.NewFuncparse(zhengai.ParseCityList,"ParseCityList") ,
	})
}


