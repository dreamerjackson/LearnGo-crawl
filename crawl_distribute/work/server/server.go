package server

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/crawl_distribute/work"
)

type CrawlService struct {

}


func (CrawlService) Process(req work.Request,result *work.ParseResult) error{

	engineReq,err:= work.DeserializeRequest(req)
	name,_:=engineReq.Parse.Serialize()
	//fmt.Printf("%s,%s\n",engineReq.Url,name)
	if err!=nil{
		return err
	}

	engineResult,err:= 	engine.Worker(engineReq)


	 *result = work.SerializeResult(engineResult)


	return nil


}