package watchdog

import "os"

func writeYAML(path string) error {
	const yamlContent = `# VM or Server (or even localhost machine)
host:
  # e.g 127.0.0.1, mydomain.com without protocol scheme like http or https
  # address: "127.0.0.1"
  address: ""

# Portainer, master address
portainer:
  # e.g 127.0.0.1, portainer.domain.com without protocol scheme like http or https
  # address: "127.0.0.1"
  address: ""
  port: 9443

# Mail account that used by docker watchdog for sending email to recipients
# e.g Developer or System Administrator
email:
  name: "Docker Watchdog"
  
  # e.g docker@isi.co.id
  sender: ""
  password: ""
  subject: "Docker Watchdog - Stopped Containers Detected (alert)"

  # Developer or System Administrator's emails
  recipients:
    - "anggit@isi.co.id"

# SMTP Server
smtp:
  address: "smtps.atisicloud.com"
  port: 587
`

	err := os.WriteFile(path, []byte(yamlContent), 0666)

	return err
}
