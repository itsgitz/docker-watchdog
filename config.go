package main

import "github.com/spf13/viper"

//Email configuration collections
type Email struct {
	Sender     string
	Password   string
	Recipient  string
	SMTPServer string
	SMTPPort   int
}

func setEmailConfig() (error, *Email) {
	//Load configuration using viper
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err, nil
	}

	//Return email configuration
	return nil, &Email{
		Sender:     viper.GetString("EMAIL_SENDER"),
		Password:   viper.GetString("EMAIL_PASSWORD"),
		Recipient:  viper.GetString("EMAIL_RECIPIENT"),
		SMTPServer: viper.GetString("SMTP_SERVER"),
		SMTPPort:   viper.GetInt("SMTP_PORT"),
	}
}
