package telebot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)
type  BotAPI struct {
	api *tgbotapi.BotAPI
}

func  NewTelegramBotAPI() (*BotAPI,error){
	botapi, err := tgbotapi.NewBotAPI("XXXX")
	if err != nil {
		return nil,err
	}
	botapi.Debug = false
	log.Printf("Authorized on account %s", botapi.Self.UserName)
	a := &BotAPI{
		api: botapi,
	}
	return a,nil
}

func (b*BotAPI) PushVideoMessage(chatid int64,file interface{}) error{
	msgvideo := tgbotapi.NewVideoUpload(chatid, file)
	_,err :=  b.api.Send(msgvideo)
	return err
}

func (b*BotAPI) PushNewMessage(chatid int64,message string) error{
	msg := tgbotapi.NewMessage(chatid, message)
	_,err :=  b.api.Send(msg)
	return err
}