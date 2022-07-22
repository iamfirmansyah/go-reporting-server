package config

import (
	"go-reporting-server/helper"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	PORT             string `mapstructure:"PORT"`
	MYSQL_CONNECTION string `mapstructure:"MYSQL_CONNECTION"`
	TELEGRAM_BOT_ID  string `mapstructure:"TELEGRAM_BOT_ID"`
	TELEGRAM_CHAT_ID string `mapstructure:"TELEGRAM_CHAT_ID"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Infof("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	helper.PanicIfError(err)

	err = viper.Unmarshal(&AppConfig)
	helper.PanicIfError(err)
}
