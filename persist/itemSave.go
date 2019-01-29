package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"crawl/LearnGo-crawl/engine"
	"errors"
)

func ItemSave() (chan engine.Item,error){
	client,err := elastic.NewClient(
		elastic.SetSniff(false))

	if err!=nil{
		return nil,err
	}

	out:=make(chan engine.Item)

	go func(){
		itemcount:=0

		for{
			item:=<-out
			log.Printf("Item saver:Got$%d,%v",itemcount,item)
			Save(client,item)
			itemcount++
		}
	}()
	return out,nil
}
func Save(client *elastic.Client,item engine.Item) error{




	if item.Type ==""{
		return errors.New("must supply Type")
	}

	indexService:= client.Index().Index("dating_profile").Type(item.Type).BodyJson(item)

	if item.Id !=""{
		indexService.Id(item.Id)
	}

	_,err  := indexService.Do(context.Background())



	if err!=nil{
		panic(err)
	}

	return nil

}


