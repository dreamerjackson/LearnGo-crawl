package parse

import (
	"crawl/LearnGo-crawl/model"
	"regexp"
	"strconv"
	"crawl/LearnGo-crawl/engine"
)





var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var public = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var  scoreRe  =  regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)

var intoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)
func ParseBookDetail(contents []byte,bookname string) engine.ParseResult{
	//fmt.Printf("%s",contents)
	bookdetail:=model.Bookdetail{}

	bookdetail.Author  =  ExtraString(contents,autoRe)

	page,err:= strconv.Atoi(ExtraString(contents,pageRe))

	if err==nil{
		bookdetail.Bookpages =  page
	}
	bookdetail.BookName = bookname
	bookdetail.Publicer = ExtraString(contents,public)
	bookdetail.Into =  ExtraString(contents,intoRe)
	bookdetail.Score = ExtraString(contents,scoreRe)
	bookdetail.Price = ExtraString(contents,priceRe)
	result:= engine.ParseResult{
		Items:[]interface{}{bookdetail},
	}

	return result
}

func ExtraString(contents []byte,re*regexp.Regexp) string{

	match:= re.FindSubmatch(contents)

	if len(match)>=2{
		return string(match[1])
	}else{
		return ""
	}
}