package main

import (
	"gopkg.in/olivere/elastic.v5"
	"crawl/LearnGo-crawl/crawl_distribute/rpcsupport"
	"crawl/LearnGo-crawl/crawl_distribute/persist"
)

func main(){

	serveRpc(":1234")
}

func serveRpc(host string) error{

	client,err:= elastic.NewClient(elastic.SetSniff(false))


	if err!=nil{

		return err
	}

	return rpcsupport.ServeRpc(host,&persist.ItemService{
		Client:client,
	})
}



