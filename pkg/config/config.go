package config

// Config - Env config
type Config struct {
	GinMode            string `yaml:"GIN_MODE"`
	BotName            string `yaml:"BOT_NAME"`
	Mid                string `yaml:"MID"`
	ChannelSecret      string `yaml:"CHANNEL_SECRET"`
	ChannelAccessToken string `yaml:"CHANNEL_ACCESS_TOKEN"`
}
