package main

import (
	"github.com/dreamerjackson/fast-crawl/engine"
	"github.com/dreamerjackson/fast-crawl/parse/ftcom"
	nytimes_politics "github.com/dreamerjackson/fast-crawl/parse/nytimes.politics"
	nytimes_science "github.com/dreamerjackson/fast-crawl/parse/nytimes.science"
	nytimes_world "github.com/dreamerjackson/fast-crawl/parse/nytimes.world"
	"github.com/dreamerjackson/fast-crawl/scheduler"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueueScheduler{},
		WorkCount:1,
	}
	seeds := []*engine.Request{
		&engine.Request{
			RootRequesrt: &engine.RootRequest{
			Url:        "https://www.ft.com/",
			ParseFunc:  ftcom.ParseHome,
			WaitPeriod:  5*time.Minute + time.Duration(rand.Int31n(5))*time.Second,
			WebFlag:"ft.com",
			},
		},
		//&engine.Request{
		//	RootRequesrt: &engine.RootRequest{
		//		Url:        "https://www.thepaper.cn/",
		//		ParseFunc:  thepaper_cn.ParseHome,
		//		WaitPeriod: 5*time.Minute + time.Duration(rand.Int31n(5))*time.Second,
		//		WebFlag:"thepaper.cn",
		//	},
		//},
		//&engine.Request{
		//	RootRequesrt: &engine.RootRequest{
		//		Url:        "https://cn.nytimes.com/",
		//		ParseFunc:  nytimes_com.ParseHome,
		//		WaitPeriod: 5*time.Minute + time.Duration(rand.Int31n(5))*time.Second,
		//		WebFlag:"nytimes.com",
		//	},
		//},
		&engine.Request{
			RootRequesrt: &engine.RootRequest{
				Url:        "https://www.nytimes.com/section/world",
				ParseFunc:  nytimes_world.ParseHome,
				WaitPeriod: 5*time.Minute + time.Duration(rand.Int31n(5))*time.Second,
				WebFlag:"nytimes_world.com",
			},
		},
		&engine.Request{
			RootRequesrt: &engine.RootRequest{
				Url:        "https://www.nytimes.com/section/politics",
				ParseFunc:  nytimes_politics.ParseHome,
				WaitPeriod: 5*time.Minute + time.Duration(rand.Int31n(5))*time.Second,
				WebFlag:"nytimes_politics.com",
			},
		},
		&engine.Request{
			RootRequesrt: &engine.RootRequest{
				Url:        "https://www.nytimes.com/section/science",
				ParseFunc:  nytimes_science.ParseHome,
				WaitPeriod: 5*time.Minute + time.Duration(rand.Int31n(5))*time.Second,
				WebFlag:"nytimes_science.com",
			},
		},
	}
	e.Start()
	e.Run(seeds...)
}