package engine

import (
	"context"
	"errors"
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/dreamerjackson/fast-crawl/fetcher"
	"github.com/dreamerjackson/fast-crawl/model/news"
	"github.com/dreamerjackson/fast-crawl/pkg/service/telebot"
	"golang.org/x/time/rate"
	"io"
	"log"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkCount int
	RootMap  map[string]*RootRequest
	//
	Api *telebot.BotAPI
	videoChan chan ParseResult
	taskChan  chan *Request
	teleChan chan string
}

type Scheduler interface {
	Submit(*Request)
	Run()
	WorkReady(chan *Request)
	WorkChan() (chan *Request)
	GetChanLen() int
}
func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}
func (e* ConcurrentEngine) Start(){
	// telegram
	e.teleChan = make(chan string,10000)
}
func(e* ConcurrentEngine) StartPush(){

	teleapi,err := telebot.NewTelegramBotAPI()
	if err !=nil{
	  fmt.Println("NewTelegramBotAPI error",err)
	}
	//
	secondLimit := rate.NewLimiter(Per(20, time.Minute), 1)
	go func() {
		for {
			select {
			case rn := <-e.teleChan:
				if err := secondLimit.Wait(context.Background());err!=nil{
					fmt.Println(err)
				}
				if err := teleapi.PushNewMessage(-1001298992352,rn);err!=nil{
					fmt.Println(err)
				}
			}
		}
	}()

}
func (e* ConcurrentEngine) Run(seeds...*Request){
	   //teleapi,err := telebot.NewTelegramBotAPI()
	   //if err !=nil{
		//   fmt.Println("NewTelegramBotAPI error",err)
	   //}
	   e.Api = nil
	   e.RootMap = make(map[string]*RootRequest,10000)
		out := make(chan ParseResult)
		e.Scheduler.Run()

		for i:=0;i<e.WorkCount;i++{
			CreateWork(e.Scheduler.WorkChan(),out,e.Scheduler)
		}
		for _,r:=range seeds{
			if r.Url == ""{
				r.ParseFunc = r.RootRequesrt.ParseFunc
				r.Url = r.RootRequesrt.Url
				r.RootRequesrt.DedupMap1 = make(map[string]struct{},100)
				r.RootRequesrt.DedupMap2 = make(map[string]struct{},100)
			}
			e.RootMap[r.RootRequesrt.WebFlag] = r.RootRequesrt
			e.Scheduler.Submit(r)
		}
	e.taskChan = make(chan *Request,len(seeds))
	// 重新开始
	go func() {
			for {
				select {
				case task := <-e.taskChan:
					go func() {
						time.Sleep(task.RootRequesrt.WaitPeriod)
						e.Scheduler.Submit(task)
					}()
				}
			}
		}()
	 // translate
    transChan := make(chan *news.News,10000)

	e.videoChan = make(chan ParseResult,10000)
	go e.PushVideo()
	go func() {
		for {
			select {
			case rn := <-transChan:
				rep,err:= gtranslate.TranslateWithParams(rn.Maintitle+" : " + rn.Subtitle, gtranslate.TranslationParams{From: "en", To: "zh"})
				if err !=nil{
					fmt.Println("err:",err)
				}
				str_time := rn.Time.Format("15:04:05")
				fmt.Println(str_time,rn.SourceName,rep)
				fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
				if rn.Push ==TELEGRAM{
					e.teleChan <- rn.SourceName+" "+rep
				}
			}
		}
	}()

		itemcount := 0
		for{
			result:=<-out
			if result.Flag == VideoFlag{
				e.videoChan<-result
				fmt.Println("123")
				continue
			}
			for _,item:= range result.Items{
				 nn,_ := item.(news.News)
				 var hashStr string = nn.Maintitle
				if len(nn.Maintitle) > 15{
					hashStr = nn.Maintitle[:15]
				}
				if _,ok:= e.RootMap[result.WebFlag].DedupMap1[hashStr];ok{
					continue
				}
				e.RootMap[result.WebFlag].DedupMap1[hashStr] = struct{}{}
				if len(e.RootMap[result.WebFlag].DedupMap1) > 9000{
						e.RootMap[result.WebFlag].DedupMap2[hashStr] = struct{}{}
				}
				if len(e.RootMap[result.WebFlag].DedupMap1) > 10000{
					e.RootMap[result.WebFlag].DedupMap1 = e.RootMap[result.WebFlag].DedupMap2
					e.RootMap[result.WebFlag].DedupMap2 = make(map[string]struct{},100)
				}
				 if nn.Type == "en"{
					 transChan<- &nn
				 }else{
				 	 str_time := nn.Time.Format("15:04:05")
					 fmt.Println(str_time,nn.SourceName,nn.Maintitle," : ",nn.Subtitle)
				 	 if nn.Push ==TELEGRAM{
						 e.teleChan <- nn.SourceName+" "+nn.Maintitle + " : " + nn.Subtitle
					 }
				 }
				//log.Printf("Got item:%d,%v",itemcount,item)
				itemcount++
			}
			for _,request:=range result.Requesrts{
				atomic.AddInt64(&request.RootRequesrt.ChildCount,1)
				e.Scheduler.Submit(request)
			}
			val := atomic.LoadInt64(&result.RootRequesrt.ChildCount)
			if val < 0 {
				e.taskChan <- &Request{
					Url:          result.RootRequesrt.Url,
					ParseFunc:    result.RootRequesrt.ParseFunc,
					RootRequesrt: result.RootRequesrt,
				}
			}
		}
}

func CreateWork(in chan *Request,out chan ParseResult,s Scheduler) {
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


func worker(r *Request) (ParseResult, error) {
	defer func() {
		atomic.AddInt64(&r.RootRequesrt.ChildCount,-1)
	}()
	//fmt.Printf("Fetch url:%s\n",r.Url)
	body,err:= fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("Fetch Error: %s",r.Url)
		return ParseResult{},err
	}
	result := r.ParseFunc(body)
	result.RootRequesrt = r.RootRequesrt
	result.WebFlag = r.RootRequesrt.WebFlag
	return result,nil
}

func (c *ConcurrentEngine) PushVideo(){
	secondLimit := rate.NewLimiter(Per(1, 15 * time.Minute), 1)
	for {
		select {
			case  result:= <-c.videoChan:
				fmt.Println("999")

				for _,item:= range result.Items {
					fmt.Println("456")

					nn, _ := item.(news.News)
					ss:= strings.Split(nn.Maintitle,"/")
					name := ss[len(ss) - 1]
					if len(ss) < 2 {
						continue
					}
					_,err:= Download(nn,name)
					if err!=nil{
						continue
					}

					if err := secondLimit.Wait(context.Background());err!=nil{
						fmt.Println(err)
					}
					if err:= c.Api.PushVideoMessage(-1001298992352,name);err !=nil{
						fmt.Println("err:",err)
					}
					err = os.Remove(name)
					if err != nil {
						fmt.Println("err:",err)
					}
				}
				c.taskChan <- &Request{
					Url:          result.RootRequesrt.Url,
					ParseFunc:    result.RootRequesrt.ParseFunc,
					RootRequesrt: result.RootRequesrt,
				}
			}
		}
}

func Download(nn news.News,name string) (*os.File,error){
	video, err := os.Create(name)
	defer video.Close()
	if err != nil  {
		return video,err
	}
	body,err:= fetcher.FetchVideo(nn.Maintitle)
	if err!=nil{
		return video,err
	}
	writen, err := io.Copy(video,body)
	if err != nil  {
		return video, err
	}
	if writen > 12 * 1024 * 1024{
		err = os.Remove(name)
		if err != nil {
			return nil,errors.New("too big")
		}
	}
	return video,nil
}