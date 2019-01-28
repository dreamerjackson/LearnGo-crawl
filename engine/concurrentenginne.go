package engine

import (
	"crawl/LearnGo-crawl/fetcher"
	"log"
)



type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan Item
}


type Scheduler interface {
	Submit(Request)
	Run()
	WorkReady(chan Request)
	WorkChan() (chan Request)
}


func (e* ConcurrentEngine) Run(seeds...Request){


		out := make(chan ParseResult)

		e.Scheduler.Run()


		for i:=0;i<e.WorkCount;i++{
			CreateWork(e.Scheduler.WorkChan(),out,e.Scheduler)
		}


		for _,r:=range seeds{
			e.Scheduler.Submit(r)
		}


		for{
			result:=<-out

			for _,item:= range result.Items{
					go func(){e.ItemChan <-item}()
			}

			for _,request:=range result.Requesrts{
				e.Scheduler.Submit(request)
			}

		}


}
func CreateWork( in chan Request,out chan ParseResult,s Scheduler) {
	go func(){
		for{

			s.WorkReady(in)
			request:= <-in

			result,err:= worker(request)

			if err!=nil{
				continue
			}
			out<-result
		}

	}()

}


func worker(r Request) (ParseResult, error) {
	//fmt.Printf("Fetch url:%s\n",r.Url)

	body,err:= fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("Fetch Error: %s",r.Url)
		return ParseResult{},err
	}

	return r.ParseFunc(body),nil
}