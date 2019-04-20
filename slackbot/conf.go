package slackbot

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	SlackConfig SlackConfig `json:"slack_config"`
}

type SlackConfig struct {
	AccessToken string `json:"access_token"`
}

func NewConfig(configPath string) (*Config, error) {
	config := new(Config)
	err := config.readConfig(configPath)
	return config, err
}

func (config *Config) readConfig(configPath string) error {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}
	return nil
}
