package dispatcher

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterDispatcher struct {
	config *oauth1.Config
	token  *oauth1.Token
	client *twitter.Client
}

func (t *TwitterDispatcher) getCredential(apiKey string, apiSecretKey string, accessToken string, accessTokenSecret string) {
	var err error

	t.config = oauth1.NewConfig(apiKey, apiSecretKey)
	t.token = oauth1.NewToken(accessToken, accessTokenSecret)

	if err != nil {
		panic(err)
	}
}

func (t *TwitterDispatcher) Start(apiKey string, apiSecretKey string, accessToken string, accessTokenSecret string) {
	t.getCredential(apiKey, apiSecretKey, accessToken, accessTokenSecret)

	httpClient := t.config.Client(oauth1.NoContext, t.token)
	t.client = twitter.NewClient(httpClient)
}

func (t *TwitterDispatcher) SendMessage(msg string) (int64, error) {
	tweet, _, err := t.client.Statuses.Update(msg, nil)

	return tweet.ID, err
}

func (t *TwitterDispatcher) ReplyMessage(msg string, statusId int64) (int64, error) {
	tweet, _, err := t.client.Statuses.Update(msg, &twitter.StatusUpdateParams{
		InReplyToStatusID: statusId,
	})

	return tweet.ID, err
}

func (t *TwitterDispatcher) DeleteEmpty(msg []string) []string {
	var r []string
	for _, str := range msg {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func (t *TwitterDispatcher) SplitMessage(msg []string) []string {
	var r []string
	for _, str := range msg {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
