package news

import "time"

type News struct {
	Maintitle string  // 主标题
	Subtitle  string  // 副标题
	Link      string  // 链接
	Type 	  string
	Push	  string // push到哪里
	SourceName   string
	Time      time.Time
}