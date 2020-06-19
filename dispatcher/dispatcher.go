package dispatcher

import (
	"bytes"
	"text/template"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
)

type Dispatcher struct {
	telegramDispatcher TelegramDispatcher
}

func (dp *Dispatcher) Start() {
	dp.telegramDispatcher.Start()
}

func (dp *Dispatcher) SendWeekly(weeklyData communityv1alpha1.WeeklySpec) {
	var tmpl string
	var output bytes.Buffer

	tmpl = "templates/Weekly.tmpl"

	tpl, err := template.ParseFiles(tmpl)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(&output, weeklyData)
	if err != nil {
		panic(err)
	}

	dp.telegramDispatcher.SendMessage(output)
	dp.telegramDispatcher.SendImage(weeklyData.ImageUrl)
}
