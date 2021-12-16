package watchdog

import (
	"errors"

	"github.com/spf13/cobra"
)

//Docker Watchdog configuration collections
type Config struct {
	Host      Host      `mapstructure:"host"`
	Portainer Portainer `mapstructure:"portainer"`
	Email     Email     `mapstructure:"email"`
	SMTP      SMTP      `mapstructure:"smtp"`
}

//Host or VM that handles the docker engine
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

func (w *Watchdog) checkConfig() {
	//Host address
	if w.Config.Host.Address == "" {
		cobra.CheckErr(
			errors.New("Host address is empty"),
		)
	}

	//Portainer address
	if w.Config.Portainer.Address == "" {
		cobra.CheckErr(
			errors.New("Portainer address is empty"),
		)
	}

	//Portainer port
	if w.Config.Portainer.Port == 0 {
		cobra.CheckErr(
			errors.New("Portainer host is empty"),
		)
	}

	//Email name
	if w.Config.Email.Name == "" {
		cobra.CheckErr(
			errors.New("Email name is empty"),
		)
	}

	//Email sender
	if w.Config.Email.Sender == "" {
		cobra.CheckErr(
			errors.New("Email sender is empty"),
		)
	}

	//Email password
	if w.Config.Email.Password == "" {
		cobra.CheckErr(
			errors.New("Email password is empty"),
		)
	}

	//Email subject
	if w.Config.Email.Subject == "" {
		cobra.CheckErr(
			errors.New("Email subject is empty"),
		)
	}

	//Email recipients
	if len(w.Config.Email.Recipients) == 0 {
		cobra.CheckErr(
			errors.New("Email recipients is empty"),
		)
	}

	//SMTP Server
	if w.Config.SMTP.Address == "" {
		cobra.CheckErr(
			errors.New("SMTP server is empty"),
		)
	}

	//SMTP port
	if w.Config.SMTP.Port == 0 {
		cobra.CheckErr(
			errors.New("SMTP port is empty"),
		)
	}
}

//CreateConfigurationFile for create new docker-watchdog configuration file
//on given path if not exist
//func CreateConfigurationFile(path, file string) {
//	//Write new configuration file
//	InformationText.Printf("[*] Create configuration file on %v \n", path)

//	viper.SafeWriteConfig()

//	filepath := path + "/" + file + ".yaml"
//	err := writeYAML(filepath)
//	if err != nil {
//		cobra.CheckErr(err)
//	}
//}
