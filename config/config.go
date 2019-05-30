package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config required by grafonnet-playground
type Config struct {
	// GrafanaURL is the url where grafana is running
	GrafanaURL string
	// GrafanaApiKey is the admin api key for grafana to create dashboards
	GrafanaAPIKey string
	// GrafonnetLibDir is the location of grafonnet lib that the app should
	// have access to. It can be managed as the situation dictates during the build
	// or package phase
	GrafonnetLibDir string
	// GrafonnetPlaygroundFolderID is the folder id in grafana where the playground
	// dashboards will be created. A separate folder is used to ensure that the
	// dashboards can be identified and deleted with ease if required
	GrafonnetPlaygroundFolderID int
}

// Load config from application.yaml file or from environment variables
func Load() *Config {
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	return &Config{
		GrafanaURL:                  mustHaveString("GRAFANA_URL"),
		GrafanaAPIKey:               mustHaveString("GRAFANA_API_KEY"),
		GrafonnetLibDir:             mustHaveString("GRAFONNET_LIB_DIR"),
		GrafonnetPlaygroundFolderID: viper.GetInt("GRAFONNET_PLAYGROUND_FOLDER_ID"),
	}
}

func mustHaveString(key string) string {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	} else {
		return viper.GetString(key)
	}
}
