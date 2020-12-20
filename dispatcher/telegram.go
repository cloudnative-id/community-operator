package dispatcher

import (
	"bytes"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type TelegramDispatcher struct {
	token      string
	credential string
	bot        *tgbotapi.BotAPI
	msg        tgbotapi.MessageConfig
	pic        tgbotapi.PhotoConfig
}

func (t *TelegramDispatcher) setCredential(token string, credential string) {
	t.token = token
	t.credential = credential
}

func (t *TelegramDispatcher) Start(token string, credential string) {
	var err error

	t.setCredential(token, credential)
	t.bot, err = tgbotapi.NewBotAPI(t.token)
	if err != nil {
		log.Fatal(err)
	}
}

func (t *TelegramDispatcher) SendMessage(output bytes.Buffer, chatType string) (int, error) {
	switch chatType {
	case "group":
		chatID, err := strconv.Atoi(t.credential)
		if err != nil {
			log.Fatal(err)
		}
		t.msg = tgbotapi.NewMessage(int64(chatID), output.String())
	case "channel":
		fmt.Println(t.credential)
		t.msg = tgbotapi.NewMessageToChannel(t.credential, output.String())
	}

	t.msg.ParseMode = "markdown"
	t.msg.DisableWebPagePreview = true

	msg, err := t.bot.Send(t.msg)
	return msg.MessageID, err
}

func (t *TelegramDispatcher) SendImage(url string, chatType string) (int, error) {
	switch chatType {
	case "group":
		chatID, err := strconv.Atoi(t.credential)
		if err != nil {
			log.Fatal(err)
		}
		t.pic = tgbotapi.NewPhotoShare(int64(chatID), url)
	case "channel":
		t.pic = tgbotapi.PhotoConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat:    tgbotapi.BaseChat{ChannelUsername: t.credential},
				FileID:      url,
				UseExisting: true,
			},
		}
	}

	msg, err := t.bot.Send(t.pic)
	return msg.MessageID, err
}
