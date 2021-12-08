package main

import (
	"github.com/spf13/viper"
)

//Docker Watchdog configuration collections
type Config struct {
	Host  Host  `mapstructure:"host"`
	Email Email `mapstructure:"email"`
	SMTP  SMTP  `mapstructure:"smtp"`
}

//Host or VM data collections
//Address and  PortainerPort
type Host struct {
	Address       string `mapstructure:"address"`
	PortainerPort int    `mapstructure:"portainer_port"`
}

type Email struct {
	Name       string   `mapstructure:"name"`
	Sender     string   `mapstructure:"sender"`
	Password   string   `mapstructure:"password"`
	Subject    string   `mapstructure:"subject"`
	Recipients []string `mapstructure:"recipients"`
}

type SMTP struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

func setConfiguration() (*Config, error) {
	//Load configuration using viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	//Unmarshalling configuration
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
