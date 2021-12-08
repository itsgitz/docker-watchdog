package main

import (
	"strings"
	"text/template"

	"github.com/docker/docker/api/types"
)

type Container struct {
	ID     string
	Name   string
	Status string
	State  string
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
			{{ range . }}
			<div>
				<p>
					<b>ID</b>: {{ .ID }}
				</p>
				<p>
					<b>Name</b>: {{ .Name }}
				</p>
				<p>
					<b>Status</b>: {{ .Status }}
				</p>
				<p>
					<b>State</b>: {{ .State }}
				</p>
			</div>
			{{ end }}
		</div>
	`

	h, err := template.New("emailbody").Parse(htmlTemplate)
	if err != nil {
		return nil, err
	}

	var containerData []Container
	containerData = setContainerDataForHTML(containers)

	err = h.Execute(&body, &containerData)
	if err != nil {
		return nil, err
	}

	bodyString := body.String()

	return &bodyString, nil
}
