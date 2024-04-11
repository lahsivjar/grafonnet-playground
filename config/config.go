package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config required by grafonnet-playground
type Config struct {
	// GrafanaPostURL is the url where grafana can be accessed for creating dashboards
	GrafanaPostURL string
	// GrafanaGetURL is the url where grafana dashboards will be loaded in iframe
	GrafanaGetURL string
	// GrafanaApiKey is the admin api key for grafana to create dashboards
	GrafanaAPIKey string
	// GrafanaAPIKeyHeaderName is the header name to sent grafanaAPIKey
	GrafanaAPIKeyHeaderName string
	// GrafonnetLibDirs is a slice of location of grafonnet lib that the app should
	// have access to. It can be managed as the situation dictates during the build
	// or package phase
	GrafonnetLibDirs []string
	// GrafonnetPlaygroundFolderID is the folder id in grafana where the playground
	// dashboards will be created. A separate folder is used to ensure that the
	// dashboards can be identified and deleted with ease if required
	GrafonnetPlaygroundFolderID int
	// AutoCleanup configures automatically cleaning up of dashboards after it is
	// loaded by the iframe
	AutoCleanup bool
	// CleanupAfter is the time after which the dashboard created is considered stale
	// and can be deleted. It doesn't stop the already loaded dashboard from working
	CleanupAfter time.Duration
	// AutoCleanupInterval configures interval between between running cleanup job
	// loaded by the iframe
	AutoCleanupInterval time.Duration
	// AutoCleanupMinBackoff configures minimum backoff for cleaner job
	AutoCleanupMinBackoff time.Duration
	// AutoCleanupMaxBackoff configures maximum backoff for cleaner job
	AutoCleanupMaxBackoff time.Duration
}

// Load config from application.yaml file or from environment variables
func Load() *Config {
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.SetDefault("AUTO_CLEANUP", "false")
	viper.SetDefault("CLEANUP_AFTER", "10s")
	viper.SetDefault("AUTO_CLEANUP_INTERVAL", "30s")
	viper.SetDefault("AUTO_CLEANUP_MIN_BACKOFF", "30s")
	viper.SetDefault("AUTO_CLEANUP_MAX_BACKOFF", "5m")
	viper.SetDefault("GRAFANA_API_KEY_HEADER_NAME", "Authorization")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	cfg := &Config{
		GrafanaPostURL:              mustHaveString("GRAFANA_POST_URL"),
		GrafanaGetURL:               mustHaveString("GRAFANA_GET_URL"),
		GrafanaAPIKey:               mustHaveString("GRAFANA_API_KEY"),
		GrafonnetLibDirs:            mustHaveStringSlice("GRAFONNET_LIB_DIRS"),
		GrafonnetPlaygroundFolderID: viper.GetInt("GRAFONNET_PLAYGROUND_FOLDER_ID"),
		GrafanaAPIKeyHeaderName:     viper.GetString("GRAFANA_API_KEY_HEADER_NAME"),
		AutoCleanup:                 viper.GetBool("AUTO_CLEANUP"),
		CleanupAfter:                viper.GetDuration("CLEANUP_AFTER"),
		AutoCleanupInterval:         viper.GetDuration("AUTO_CLEANUP_INTERVAL"),
		AutoCleanupMinBackoff:       viper.GetDuration("AUTO_CLEANUP_MIN_BACKOFF"),
		AutoCleanupMaxBackoff:       viper.GetDuration("AUTO_CLEANUP_MAX_BACKOFF"),
	}
	return cfg
}

func mustHaveString(key string) string {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	} else {
		return viper.GetString(key)
	}
}

func mustHaveStringSlice(key string) []string {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	} else {
		return viper.GetStringSlice(key)
	}
}
