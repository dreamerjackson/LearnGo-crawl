package engine

import (
	"crawl/LearnGo-crawl/fetcher"
	"log"
)

type simleEngine struct{

}

func (s simleEngine)Run(seeds...Request){

	var requests []Request

	for _,e:=range seeds{
		requests = append(requests,e)
	}

	for len(requests)>0{

		r:=requests[0]

		requests = requests[1:]
		log.Printf("Fetching url:%s",r.Url)
		body,err:= fetcher.Fetch(r.Url)

		if err!=nil{
			log.Printf("Fetch Error: %s",r.Url)
		}

		parseresult:=  r.Parse.Parse(body,r.Url)

		requests = append(requests,parseresult.Requesrts...)

		for _,item:=range parseresult.Items{
			log.Printf("Got item:%s",item)
		}
	}
}