package client

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/crawl_distribute/rpcsupport"
	"crawl/LearnGo-crawl/crawl_distribute/work"
)

func CreateProcess() (engine.Processor,error){
		client,err:= rpcsupport.NewClient(":1235")

		if err!=nil{
			return nil,err
		}

		return func(req engine.Request)(engine.ParseResult,error){
			sReq:= work.SerializeRequest(req)

			var sResult work.ParseResult

			err:= client.Call("CrawlService.Process",sReq,&sResult)

			if err!=nil{
				return engine.ParseResult{},nil
			}
			return work.DeserializeResult(sResult),nil

		},nil

}