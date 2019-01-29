package zhengai

import (
	"regexp"
	"crawl/LearnGo-crawl/engine"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)


func ParseCity(contents []byte,_ string) engine.ParseResult{

	matches:= cityRe.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	for _,m:= range matches{

		//url:=string(m[1])
		name:=string(m[2])
		//println(string(m[1]))
		//不用用户名了
		//result.Items = append(result.Items,"User:"+string(m[2]))
		result.Requesrts = append(result.Requesrts,engine.Request{
			Url:string(m[1]),
			Parse:NewprofileParse(name),
		})
	}

	//查找城市页面下的城市链接
	matches=  cityUrlRe.FindAllSubmatch(contents,-1)

	for _,m:= range matches{
		result.Requesrts = append(result.Requesrts,engine.Request{
			Url:string(m[1]),
			Parse:engine.NewFuncparse(ParseCity,"Parsecity"),
		})
	}


	return result

}
