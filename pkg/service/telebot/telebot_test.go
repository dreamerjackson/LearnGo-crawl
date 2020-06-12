package telebot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBotAPI_PushVideoMessage(t *testing.T) {
	api,err:= NewTelegramBotAPI()
	assert.NoError(t,err)
	err = api.PushVideoMessage(-1001298992352,"/Users/jackson/Desktop/SilverFatherlyCanary.mp4")
	assert.NoError(t,err)
}