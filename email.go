package main

import (
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"gopkg.in/gomail.v2"
)

func sendEmailAlert(containers []types.Container) {
	config, err := setEmailConfig()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(config.Sender, config.Name)},
		"To":      config.Recipients,
		"Subject": {config.Subject},
	})

	//Set HTML format for email body
	emailBody := setEmailBody(containers, config.Host.Address)

	m.SetBody("text/html", string(emailBody))

	d := gomail.NewDialer(config.SMTP.Address, config.SMTP.Port, config.Sender, config.Password)

	//Send email to recipient
	informationText.Printf("[*] Sending email alert to: %v \n", config.Recipients)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func setEmailBody(containers []types.Container, host string) []byte {
	var htmlBody []byte

	containerHost := fmt.Sprintf("<p><b>Docker Watchdog</b> has detected stopped containers on %v!</p>", host)
	htmlBody = append(htmlBody, containerHost...)

	for _, c := range containers {

		htmlBody = append(htmlBody, "<br>"...)
		htmlBody = append(htmlBody, "<div>"...)

		//Write container ID
		containerID := fmt.Sprintf("<p><b>ID</b>: %v</p>", c.ID[:10])
		htmlBody = append(htmlBody, containerID...)

		//Write container Name
		containerName := fmt.Sprintf("<p><b>Name</b>: %v</p>", strings.Trim(c.Names[0], "/"))
		htmlBody = append(htmlBody, containerName...)

		//Write container Status
		containerStatus := fmt.Sprintf("<p><b>Status</b>: %v</p>", c.Status)
		htmlBody = append(htmlBody, containerStatus...)

		//Write container State
		containerState := fmt.Sprintf("<p><b>State</b>: %v</p>", c.State)
		htmlBody = append(htmlBody, containerState...)

		htmlBody = append(htmlBody, "</div>"...)
		htmlBody = append(htmlBody, "<br>"...)
	}

	return htmlBody
}
