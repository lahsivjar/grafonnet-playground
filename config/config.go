package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	GrafanaUrl    string
	GrafanaApiKey string
}

func Load() *Config {
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	return &Config{
		GrafanaUrl:    mustHaveString("GRAFANA_URL"),
		GrafanaApiKey: mustHaveString("GRAFANA_API_KEY"),
	}
}

func mustHaveString(key string) string {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	} else {
		return viper.GetString(key)
	}
}
