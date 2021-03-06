package watchdog

import (
	"fmt"
	"strings"
)

//setPortainerAddress set the portainer address as a combination of
//`host` and `portainer_port`
func setPortainerAddress(config *Config) string {
	//Use strings.Builder for fastest string concatenation
	var addr strings.Builder

	fmt.Fprintf(&addr, "https://%s:%d", config.Portainer.Address, config.Portainer.Port)

	return addr.String()
}
