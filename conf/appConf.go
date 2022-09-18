package conf

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

var LOG, _ = zap.NewDevelopment()

func LoadAppConfiguration() error {
	viper.AddConfigPath("./resources")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		LOG.Sugar().Info("configuration file changed: %s", e.Name)
	})

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error while loading configuration from file %w", err)
	}

	err = loadConfigurationFormServer(viper.GetString("dev.config-server-url"))
	if err != nil {
		return fmt.Errorf("error while loading configuration %s", err)
	}

	return nil
}

func loadConfigurationFormServer(configServerUrl string) error {
	resp, err := http.Get(configServerUrl)
	if err != nil {
		return fmt.Errorf("could not load configuration error: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error while parsing response body: %w", err)
	}

	err = parseConfigurationBody(body)
	if err != nil {
		return fmt.Errorf("error while parsing configuration: %w", err)
	}

	return nil
}

func parseConfigurationBody(body []byte) error {
	var cc CloudConfig

	err := json.Unmarshal(body, &cc)
	if err != nil {
		return fmt.Errorf("error while parsing response body: %w", err)
	}

	for key, value := range cc.PropertySource[0].Source {
		viper.Set(key, value)
		fmt.Printf("loading config property %s => %v\n", key, value)
	}

	if viper.IsSet("server_name") {
		fmt.Printf("successfully loaded configuration for service %s\n", viper.GetString("server_name"))
	}

	return nil
}

type CloudConfig struct {
	Name           string           `json:"name"`
	Profiles       []string         `json:"profiles"`
	Label          string           `json:"label"`
	Version        string           `json:"version"`
	PropertySource []PropertySource `json:"propertySources"`
}

type PropertySource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}
