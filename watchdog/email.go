package watchdog

import (
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
	"gopkg.in/gomail.v2"
)

const htmlContentType = "text/html"

func (w *Watchdog) sendEmailAlert(containers []types.Container) {
	//Load initial configuration
	emailCached, err := isCached(containers, w.Config)
	if err != nil {
		panic(err)
	}

	//Send email alert if cache doesn't exist
	if !emailCached {
		//Initialize gomail email header and body
		m := gomail.NewMessage()
		m.SetHeaders(map[string][]string{
			"From": {
				m.FormatAddress(w.Config.Email.Sender, w.Config.Email.Name),
			},
			"To":      w.Config.Email.Recipients,
			"Subject": {w.Config.Email.Subject},
		})

		//Set HTML format for email body
		emailBody, err := setEmailBody(containers, w.Config)
		if err != nil {
			cobra.CheckErr(err)
		}
		m.SetBody(htmlContentType, *emailBody)

		//Define gomail email client
		d := gomail.NewDialer(
			w.Config.SMTP.Address,
			w.Config.SMTP.Port,
			w.Config.Email.Sender,
			w.Config.Email.Password,
		)

		//Send email to recipients
		InformationText.Printf(
			"[*] Sending email alert to: %v \n",
			w.Config.Email.Recipients,
		)

		if err := d.DialAndSend(m); err != nil {
			cobra.CheckErr(err)
		}
	}
}

func setEmailBody(containers []types.Container, config *Config) (*string, error) {
	htmlBody, err := setHTMLBody(containers, config)

	return htmlBody, err
}
