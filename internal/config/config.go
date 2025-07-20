package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	TelegramBotToken string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	config := &Config{
		TelegramBotToken: viper.GetString("TELEGRAM_BOT_TOKEN"),
	}

	if config.TelegramBotToken == "" {
		return nil, errors.New("telegram bot token not set")
	}
	
	return config, nil
}
