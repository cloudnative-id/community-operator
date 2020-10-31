package dispatcher

import (
	"bytes"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramDispatcher struct {
	token  string
	chatID int
	bot    *tgbotapi.BotAPI
	msg    tgbotapi.MessageConfig
	pic    tgbotapi.PhotoConfig
}

func (t *TelegramDispatcher) setCredential(token string, chatID string) {
	var err error

	t.token = token
	t.chatID, err = strconv.Atoi(chatID)
	if err != nil {
		panic(err)
	}
}

func (t *TelegramDispatcher) Start(token string, chatID string) {
	var err error

	t.setCredential(token, chatID)
	t.bot, err = tgbotapi.NewBotAPI(t.token)
	if err != nil {
		panic(err)
	}
}

func (t *TelegramDispatcher) SendMessage(output bytes.Buffer) (int, error) {
	t.msg = tgbotapi.NewMessage(int64(t.chatID), output.String())
	t.msg.ParseMode = "markdown"
	t.msg.DisableWebPagePreview = true

	msg, err := t.bot.Send(t.msg)
	return msg.MessageID, err
}

func (t *TelegramDispatcher) SendImage(url string) (int, error) {
	t.pic = tgbotapi.NewPhotoShare(int64(t.chatID), url)

	msg, err := t.bot.Send(t.pic)
	return msg.MessageID, err
}
