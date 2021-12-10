package main

import (
	"github.com/docker/docker/api/types"
	"gopkg.in/gomail.v2"
)

const htmlContentType = "text/html"

func sendEmailAlert(containers []types.Container) {
	//Load initial configuration
	config, err := setConfiguration()
	if err != nil {
		panic(err)
	}

	emailCached, err := isCached(containers)
	if err != nil {
		panic(err)
	}

	//Send email alert if cache doesn't exist
	if !emailCached {
		//Initialize gomail email header and body
		m := gomail.NewMessage()
		m.SetHeaders(map[string][]string{
			"From":    {m.FormatAddress(config.Email.Sender, config.Email.Name)},
			"To":      config.Email.Recipients,
			"Subject": {config.Email.Subject},
		})

		//Set HTML format for email body
		emailBody, err := setEmailBody(containers, config)
		if err != nil {
			panic(err)
		}
		m.SetBody(htmlContentType, *emailBody)

		//Define gomail email client
		d := gomail.NewDialer(config.SMTP.Address, config.SMTP.Port, config.Email.Sender, config.Email.Password)

		//Send email to recipients
		informationText.Printf("[*] Sending email alert to: %v \n", config.Email.Recipients)
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}

func setEmailBody(containers []types.Container, config *Config) (*string, error) {
	htmlBody, err := setHTMLBody(containers, config)

	return htmlBody, err
}
