package dispatcher

import (
	"bytes"
	"os"
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

func (t *TelegramDispatcher) getCredential() {
	var err error

	t.token = os.Getenv("TELEGRAM_TOKEN")
	t.chatID, err = strconv.Atoi(os.Getenv("TELEGRAM_CHATID"))
	if err != nil {
		panic(err)
	}
}

func (t *TelegramDispatcher) Start() {
	var err error

	t.getCredential()
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
