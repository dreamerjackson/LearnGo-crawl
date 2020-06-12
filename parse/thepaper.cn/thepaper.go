package thepaper_cn

import (
	"github.com/dreamerjackson/fast-crawl/engine"
	"github.com/dreamerjackson/fast-crawl/model/news"
	"regexp"
	"strings"
	"time"
)

var headerRe = regexp.MustCompile(`<div class="news_li"[\s\S]*?<h2>[\s\S]*?<a.*?target="_blank">([\s\S]*?)</a>[\s\S]*?<p>([\s\S]*?)</p>`)

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
			Subtitle:  strings.TrimSpace(strings.ReplaceAll(string(m[2]),"\n","")),
			Time:now,
			SourceName: "澎湃新闻",
			Push: engine.TELEGRAM,
			// Link:      "",
		}
		//不用用户名了
		result.Items = append(result.Items,item)
	}
	return result
}
