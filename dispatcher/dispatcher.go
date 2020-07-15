package dispatcher

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
)

type Dispatcher struct {
	telegramEnabled    bool
	twitterEnabled     bool
	telegramDispatcher TelegramDispatcher
	twitterDispacher   TwitterDispatcher
}

func (dp *Dispatcher) getCredential() {
	var err error

	dp.telegramEnabled, err = strconv.ParseBool(os.Getenv("TELEGRAM_ENABLED"))
	if err != nil {
		panic(err)
	}

	dp.twitterEnabled, err = strconv.ParseBool(os.Getenv("TWITTER_ENABLED"))
	if err != nil {
		panic(err)
	}
}

func (dp *Dispatcher) Start() {
	dp.getCredential()

	if dp.telegramEnabled {
		dp.telegramDispatcher.Start()
	}

	if dp.twitterEnabled {
		dp.twitterDispacher.Start()
	}
}

func (dp *Dispatcher) sendTelegram(message bytes.Buffer, picture string) {
	_, err := dp.telegramDispatcher.SendMessage(message)
	if err != nil {
		panic(err)
	}

	_, err = dp.telegramDispatcher.SendImage(picture)
	if err != nil {
		panic(err)
	}
}

func (dp *Dispatcher) sendTwitter(message bytes.Buffer, picture string) {
	temp := message.String()
	msg := dp.twitterDispacher.SplitMessage(dp.twitterDispacher.DeleteEmpty(strings.Split(temp, "\n")))

	id, err := dp.twitterDispacher.SendMessage(msg[0])
	if err != nil {
		panic(err)
	}

	for i := 1; i < len(msg); i++ {
		id, err = dp.twitterDispacher.ReplyMessage(msg[i], id)
		if err != nil {
			panic(err)
		}
	}
}

func format(dispatcher string, data interface{}) bytes.Buffer {
	var tmpl string
	var output bytes.Buffer

	switch data.(type) {
	case communityv1alpha1.WeeklySpec:
		tmpl = fmt.Sprintf("templates/%s/Weekly.tmpl", dispatcher)
	case communityv1alpha1.MeetupSpec:
		tmpl = fmt.Sprintf("templates/%s/Meetup.tmpl", dispatcher)
	case communityv1alpha1.AnnouncementSpec:
		tmpl = fmt.Sprintf("templates/%s/Announcement.tmpl", dispatcher)
	}

	tpl, err := template.ParseFiles(tmpl)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(&output, data)
	if err != nil {
		panic(err)
	}

	return output
}

func (dp *Dispatcher) SendWeekly(weeklyData communityv1alpha1.WeeklySpec) {

	sort.Slice(weeklyData.Articles, func(i, j int) bool {
		return weeklyData.Articles[i].Type < weeklyData.Articles[j].Type
	})

	telegramData := format("telegram", weeklyData)
	twitterData := format("twitter", weeklyData)

	if dp.telegramEnabled {
		dp.sendTelegram(telegramData, weeklyData.Image)
	}

	if dp.twitterEnabled {
		dp.sendTwitter(twitterData, weeklyData.Image)
	}
}

func (dp *Dispatcher) SendMeetup(meetupData communityv1alpha1.MeetupSpec) {
	telegramData := format("telegram", meetupData)
	twitterData := format("twitter", meetupData)

	if dp.telegramEnabled {
		dp.sendTelegram(telegramData, meetupData.Image)
	}

	if dp.twitterEnabled {
		dp.sendTwitter(twitterData, meetupData.Image)
	}
}

func (dp *Dispatcher) SendAnnouncement(announcementData communityv1alpha1.AnnouncementSpec) {
	telegramData := format("telegram", announcementData)
	twitterData := format("twitter", announcementData)

	if dp.telegramEnabled {
		dp.sendTelegram(telegramData, announcementData.Image)
	}

	if dp.twitterEnabled {
		dp.sendTwitter(twitterData, announcementData.Image)
	}
}
