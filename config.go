package main

import (
	"github.com/spf13/viper"
)

//Docker Watchdog configuration collections
type Config struct {
	Host      Host      `mapstructure:"host"`
	Portainer Portainer `mapstructure:"portainer"`
	Email     Email     `mapstructure:"email"`
	SMTP      SMTP      `mapstructure:"smtp"`
}

//Host or VM data collections
//Address
type Host struct {
	Address string `mapstructure:"address"`
}

//Portainer address and port
type Portainer struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

//Email configuration
type Email struct {
	Name       string   `mapstructure:"name"`
	Sender     string   `mapstructure:"sender"`
	Password   string   `mapstructure:"password"`
	Subject    string   `mapstructure:"subject"`
	Recipients []string `mapstructure:"recipients"`
}

//SMTP or mail server agent configuration
type SMTP struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

func setConfiguration() (*Config, error) {
	//Load configuration using viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	//Viper is reading the configuration here ...
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
