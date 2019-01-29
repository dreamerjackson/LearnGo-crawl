package main

import (
	"log"
	"crawl/LearnGo-crawl/crawl_distribute/rpcsupport"
	"crawl/LearnGo-crawl/crawl_distribute/work/server"
)

func main(){
	log.Fatal(rpcsupport.ServeRpc(":1235",&server.CrawlService{}))
}
