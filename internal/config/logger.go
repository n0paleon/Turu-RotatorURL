package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(config *viper.Viper) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(config.GetInt32("app.log_level")))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
