package config

type Config struct {
	Telegram struct {
		Enabled bool `yaml:"enabled"`
		Group   []struct {
			ChatID string `yaml:"chatId"`
			Token  string `yaml:"token"`
		} `yaml:"group"`
		Channel []struct {
			Username string `yaml:"username"`
			Token    string `yaml:"token"`
		} `yaml:"channel"`
	} `yaml:"telegram"`
	Twitter struct {
		Enabled bool `yaml:"enabled"`
		Account []struct {
			APIKey            string `yaml:"apiKey"`
			APISecretKey      string `yaml:"apiSecretKey"`
			AccessToken       string `yaml:"accessToken"`
			AccessTokenSecret string `yaml:"accessTokenSecret"`
		} `yaml:"account"`
	} `yaml:"twitter"`
}
