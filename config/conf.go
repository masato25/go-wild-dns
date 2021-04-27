package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func ReadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}
	fillEnvWithConf()
	loggerSetup()
}

func fillEnvWithConf() {
	if os.Getenv("DOMAIN_SUFFIX") == "" {
		os.Setenv("DOMAIN_SUFFIX", viper.GetString("env.DOMAIN_SUFFIX"))
	}
}

func loggerSetup() {
	switch viper.Get("logger.level") {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
