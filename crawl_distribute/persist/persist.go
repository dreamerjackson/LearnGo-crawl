package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/persist"
)

type ItemService struct{
	Client * elastic.Client
}

func (s*ItemService) Save(item engine.Item,result*string) error{
	err := persist.Save(s.Client,item)

	if err==nil{
		*result = "ok"
	}

	return err

}