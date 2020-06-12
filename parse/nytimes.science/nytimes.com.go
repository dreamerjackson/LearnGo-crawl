package nytimes_world

import (
	"github.com/dreamerjackson/fast-crawl/engine"
	"github.com/dreamerjackson/fast-crawl/model/news"
	"regexp"
	"time"
)

var headerRe = regexp.MustCompile(`<h2 class="css-1j9dxys e1xfvim30">(.*?)</h2><p class="css-1echdzn e1xfvim31">(.*?)</p>`)
func ParseHome(contents []byte) engine.ParseResult{
	matches:= headerRe.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	result.Flag = engine.FlagEnd
	result.Items = make([]interface{},0,5)
	now := time.Now()
	for _,m:= range matches{
		if len(m) < 3 {
			continue
		}
		item := news.News{
			Maintitle: string(m[1]),
			Subtitle:  string(m[2]),
			Type:"en",
			Time:now,
			SourceName:"纽时科学",
			Push: engine.TELEGRAM,
			// Link:      "",
		}
		//不用用户名了
		result.Items = append(result.Items,item)
	}
	return result

}
