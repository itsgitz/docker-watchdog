package main

import (
	"strings"
	"text/template"

	"github.com/docker/docker/api/types"
)

type HTMLDataBind struct {
	HostAddress      string
	PortainerAddress string
	Container        []Container
}

func setContainerDataForHTML(containers []types.Container) []Container {
	var container []Container

	for _, c := range containers {
		container = append(container, Container{
			ID:     c.ID[:10],
			Name:   strings.Trim(c.Names[0], "/"),
			Status: c.Status,
			State:  c.State,
		})
	}

	return container
}

func setHTMLBody(containers []types.Container, config *Config) (*string, error) {
	var body strings.Builder

	const htmlTemplate = `
		<div>
			<b>Docker Watchdog</b>
			has detected stopped containers on your server ({{ .HostAddress }})
			<div style="padding: 3px;"></div>
		</div>
		<div>
			<table border="1" cellpadding="10" cellspacing="0">
				<thead>
					<tr>
						<th><b>ID</b></th>
						<th><b>Name</b></th>
						<th><b>Status</b></th>
						<th><b>State</b></th>
					</tr>
				</thead>	
				{{ range .Container }}
				<tr>
					<td>{{ .ID }}</td>
					<td>{{ .Name }}</td>
					<td>{{ .Status }}</td>
					<td>
						<span style="color: red;">
							<b>{{ .State }}</b>
						</span>
					</td>
				</tr>
				{{ end }}
			</table>	
		</div>
		<div>
			<div style="padding: 3px;"></div>
			<a href="{{ .PortainerAddress }}" target="__blank">
				Check on Portainer
			</a>
		</div>
	`

	h, err := template.New("emailbody").Parse(htmlTemplate)
	if err != nil {
		return nil, err
	}

	//Set container data for bind with html
	var containerData []Container
	containerData = setContainerDataForHTML(containers)

	//Set portainer address
	var portainerAddress = setPortainerAddress(config)

	//Populate the HTML data bind
	var htmlDataBind HTMLDataBind = HTMLDataBind{
		HostAddress:      config.Host.Address,
		PortainerAddress: portainerAddress,
		Container:        containerData,
	}

	err = h.Execute(&body, &htmlDataBind)
	if err != nil {
		return nil, err
	}

	bodyString := body.String()

	return &bodyString, nil
}
