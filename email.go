package main

import "gopkg.in/gomail.v2"

func sendEmailAlert() {
	err, config := setEmailConfig()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.Sender)
	m.SetHeader("To", config.Recipient)
	m.SetHeader("Subject", "Docker Watchdog - Stopped Containers Detected")
	m.SetBody("text/html", "There are stopped containers detected!")

	d := gomail.NewDialer(config.SMTPServer, config.SMTPPort, config.Sender, config.Password)

	//Send email to recipient
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
