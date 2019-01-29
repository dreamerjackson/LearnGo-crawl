package zhengai

import (
	"regexp"
	"crawl/LearnGo-crawl/engine"
)

const cityListRe = `(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte,_ string) engine.ParseResult{
	re:=regexp.MustCompile(cityListRe)

	matches:= re.FindAllSubmatch(contents,-1)
	result:=engine.ParseResult{}

	for _,m :=range matches{
		//result.Items = append(result.Items,string(m[2]))
		result.Requesrts  = append(
			result.Requesrts,engine.Request{
				Url:string(m[1]),
				Parse:engine.NewFuncparse(ParseCity,"Parsecity"),
			})
	}
	return result
}
