package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type AppConfigAuthRegistration struct {
	Enabled bool `yaml:"enabled"`
}

type AppConfigAuthGoogle struct {
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type AppConfigAuth struct {
	Strategies   []string                  `yaml:"strategies"`
	Registration AppConfigAuthRegistration `yaml:"registration"`
	Google       AppConfigAuthGoogle       `yaml:"google"`
}

type AppConfig struct {
	Auth AppConfigAuth `yaml:"auth"`
}

var (
	Config *AppConfig
)

func init() {

	configFilePath := os.Getenv("FINCAL_CFG")

	if configFilePath == "" {
		configFilePath = "config.yaml"
	}

	config := AppConfig{
		AppConfigAuth{
			Strategies:   []string{"credentials"},
			Registration: AppConfigAuthRegistration{Enabled: true},
			Google: AppConfigAuthGoogle{
				ClientId:     "abx",
				ClientSecret: "xyz",
			},
		},
	}

	data, err := os.ReadFile(configFilePath)

	if err != nil {
		log.Fatalf("Failed to open config file: %s", err)
	}

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		log.Fatalf("Failed to parse config file: %s", err)
	}

	Config = &config
}
