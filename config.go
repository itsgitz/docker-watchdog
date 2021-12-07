package main

import (
	"github.com/spf13/viper"
)

//Email configuration collections

type Config struct {
	Host  Host  `mapstructure:"host"`
	Email Email `mapstructure:"email"`
	SMTP  SMTP  `mapstructure:"smtp"`
}

type Host struct {
	Address string
}

type Email struct {
	Name       string
	Sender     string
	Password   string
	Subject    string
	Recipients []string
	SMTP       SMTP
	Host       Host
}

type SMTP struct {
	Address string
	Port    int
}

func setEmailConfig() (*Email, error) {
	//Load configuration using viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Email{
		Name:       viper.GetString("email.name"),
		Sender:     viper.GetString("email.sender"),
		Password:   viper.GetString("email.password"),
		Subject:    viper.GetString("email.subject"),
		Recipients: viper.GetStringSlice("email.recipients"),
		SMTP: SMTP{
			Address: viper.GetString("smtp.address"),
			Port:    viper.GetInt("smtp.port"),
		},
		Host: Host{
			Address: viper.GetString("host.address"),
		},
	}, nil
}
