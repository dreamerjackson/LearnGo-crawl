package main

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/parse"
)

func main(){
	engine.Run(engine.Request{
		Url:"https://book.douban.com",
		ParseFunc:parse.ParseContent,
	})

}




