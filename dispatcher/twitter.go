package dispatcher

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterDispatcher struct {
	config *oauth1.Config
	token  *oauth1.Token
	client *twitter.Client
}

func (t *TwitterDispatcher) getCredential() {
	var err error

	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecretKey := os.Getenv("TWITTER_API_SECRET_KEY")
	t.config = oauth1.NewConfig(apiKey, apiSecretKey)

	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	t.token = oauth1.NewToken(accessToken, accessTokenSecret)

	if err != nil {
		panic(err)
	}
}

func (t *TwitterDispatcher) Start() {
	t.getCredential()

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
