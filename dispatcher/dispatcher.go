package dispatcher

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/cloudnative-id/community-operator/config"
	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"
)

type Dispatcher struct {
	config             config.Config
	telegramDispatcher TelegramDispatcher
	twitterDispacher   TwitterDispatcher
}

func (dp *Dispatcher) getConfig() {
	configFile, err := ioutil.ReadFile("/etc/community-operator/config/community-operator-config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configFile, &dp.config)
	if err != nil {
		log.Fatal(err)
	}
}

func (dp *Dispatcher) Start() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	dp.getConfig()
}

func (dp *Dispatcher) sendTelegram(message bytes.Buffer, picture string, chatType string) {
	_, err := dp.telegramDispatcher.SendMessage(message, chatType)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	_, err = dp.telegramDispatcher.SendImage(picture, chatType)
	if err != nil {
		log.Error(err)
		log.Error("Cannot send image, skipping")
	}
}

func (dp *Dispatcher) sendTwitter(message bytes.Buffer, picture string) {
	temp := message.String()
	msg := dp.twitterDispacher.SplitMessage(dp.twitterDispacher.DeleteEmpty(strings.Split(temp, "\n")))

	id, err := dp.twitterDispacher.SendMessage(msg[0])
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	for i := 1; i < len(msg); i++ {
		id, err = dp.twitterDispacher.ReplyMessage(msg[i], id)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(5 * time.Second)
	}
}

func populateTemplate(dispatcher string, data interface{}) bytes.Buffer {
	var tmpl string
	var output bytes.Buffer

	switch data.(type) {
	case communityv1alpha1.WeeklySpec:
		tmpl = fmt.Sprintf("/etc/community-operator/templates/%s/Weekly.tmpl", dispatcher)
	case communityv1alpha1.MeetupSpec:
		tmpl = fmt.Sprintf("/etc/community-operator/templates/%s/Meetup.tmpl", dispatcher)
	case communityv1alpha1.AnnouncementSpec:
		tmpl = fmt.Sprintf("/etc/community-operator/templates/%s/Announcement.tmpl", dispatcher)
	}

	tpl, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(&output, data)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

func (dp *Dispatcher) SendWeekly(weeklyData communityv1alpha1.WeeklySpec) {
	sort.Slice(weeklyData.Articles, func(i, j int) bool {
		return weeklyData.Articles[i].Type < weeklyData.Articles[j].Type
	})

	if dp.config.Telegram.Enabled {
		for _, config := range dp.config.Telegram.Group {
			dp.telegramDispatcher.Start(config.Token, config.ChatID)

			telegramData := populateTemplate("telegram", weeklyData)
			dp.sendTelegram(telegramData, weeklyData.Image, "group")
		}

		for _, config := range dp.config.Telegram.Channel {
			dp.telegramDispatcher.Start(config.Token, config.Username)

			telegramData := populateTemplate("telegram", weeklyData)
			dp.sendTelegram(telegramData, weeklyData.Image, "channel")
		}
	}

	if dp.config.Twitter.Enabled {
		for _, config := range dp.config.Twitter.Account {
			dp.twitterDispacher.Start(config.APIKey, config.APISecretKey, config.AccessToken, config.AccessTokenSecret)

			twitterData := populateTemplate("twitter", weeklyData)
			dp.sendTwitter(twitterData, weeklyData.Image)
		}
	}
}

func (dp *Dispatcher) SendMeetup(meetupData communityv1alpha1.MeetupSpec) {
	if dp.config.Telegram.Enabled {
		for _, config := range dp.config.Telegram.Group {
			dp.telegramDispatcher.Start(config.Token, config.ChatID)

			telegramData := populateTemplate("telegram", meetupData)
			dp.sendTelegram(telegramData, meetupData.Image, "group")
		}

		for _, config := range dp.config.Telegram.Channel {
			dp.telegramDispatcher.Start(config.Token, config.Username)

			telegramData := populateTemplate("telegram", meetupData)
			dp.sendTelegram(telegramData, meetupData.Image, "channel")
		}
	}

	if dp.config.Twitter.Enabled {
		for _, config := range dp.config.Twitter.Account {
			dp.twitterDispacher.Start(config.APIKey, config.APISecretKey, config.AccessToken, config.AccessTokenSecret)

			twitterData := populateTemplate("twitter", meetupData)
			dp.sendTwitter(twitterData, meetupData.Image)
		}
	}
}

func (dp *Dispatcher) SendAnnouncement(announcementData communityv1alpha1.AnnouncementSpec) {
	if dp.config.Telegram.Enabled {
		for _, config := range dp.config.Telegram.Group {
			dp.telegramDispatcher.Start(config.Token, config.ChatID)

			telegramData := populateTemplate("telegram", announcementData)
			dp.sendTelegram(telegramData, announcementData.Image, "group")
		}

		for _, config := range dp.config.Telegram.Channel {
			dp.telegramDispatcher.Start(config.Token, config.Username)

			telegramData := populateTemplate("telegram", announcementData)
			dp.sendTelegram(telegramData, announcementData.Image, "channel")
		}
	}

	if dp.config.Twitter.Enabled {
		for _, config := range dp.config.Twitter.Account {
			dp.twitterDispacher.Start(config.APIKey, config.APISecretKey, config.AccessToken, config.AccessTokenSecret)

			twitterData := populateTemplate("twitter", announcementData)
			dp.sendTwitter(twitterData, announcementData.Image)
		}
	}
}
