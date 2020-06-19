package dispatcher

import (
	"bytes"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramDispatcher struct {
	Token  string
	ChatID int
	Bot    *tgbotapi.BotAPI
	Msg    tgbotapi.MessageConfig
	Pic    tgbotapi.PhotoConfig
}

func (t *TelegramDispatcher) GetCredential() {
	var err error

	t.Token = os.Getenv("TELEGRAM_TOKEN")
	t.ChatID, err = strconv.Atoi(os.Getenv("TELEGRAM_CHATID"))
	if err != nil {
		panic(err)
	}
}

func (t *TelegramDispatcher) Start() {
	var err error

	t.GetCredential()
	t.Bot, err = tgbotapi.NewBotAPI(t.Token)
	if err != nil {
		panic(err)
	}
}

func (t *TelegramDispatcher) SendMessage(output bytes.Buffer) {
	t.Msg = tgbotapi.NewMessage(int64(t.ChatID), output.String())
	t.Msg.ParseMode = "markdown"
	t.Msg.DisableWebPagePreview = true
	t.Bot.Send(t.Msg)
}

func (t *TelegramDispatcher) SendImage(url string) {
	t.Pic = tgbotapi.NewPhotoShare(int64(t.ChatID), url)
	t.Bot.Send(t.Pic)
}
