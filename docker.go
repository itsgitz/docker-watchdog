package main

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

//Define stopped containers state
//status=(paused|exited|dead)
//src: https://docs.docker.com/engine/api/v1.41/#operation/ContainerList
const (
	ExitedState = "exited" //Exited
	PausedState = "paused" //Paused
	DeadState   = "dead"   //Dead
)

type Container struct {
	ID     string
	Name   string
	Status string
	State  string
}

func getStoppedContainers() {
	ctx := context.Background()

	//Create new API client
	//WithAPIVersionNegotiation enables automatic API version negotiation for the client
	//With this option enabled, the client automatically negotiates the API version to use when making requests
	//src: https://pkg.go.dev/github.com/docker/docker/client#WithAPIVersionNegotiation
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	//Get all containers
	//Filters by exited, paused, or dead status
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(
			filters.Arg("status", ExitedState), //Exited
			filters.Arg("status", PausedState), //Paused
			filters.Arg("status", DeadState),   //Dead
		),
	})
	if err != nil {
		panic(err)
	}

	informationText.Printf("[*] Searching for stopped containers (exited|paused|dead) ... \n")

	//Return nothing if there are not stopped containers
	//And continue detection ...
	if len(containers) == 0 {
		successText.Printf("[*] All containers are running ... \n\n")
		return
	}

	//Display all stopped containers information
	dangerText.Printf("\n[!] Stopped container is detected! \n\n")
	for _, c := range containers {
		dangerText.Printf("| ID: %v \n", c.ID[:10])
		dangerText.Printf("| Name: %v \n", strings.Trim(c.Names[0], "/"))
		dangerText.Printf("| Status: %v \n", c.Status)
		dangerText.Printf("| State: %v \n\n", c.State)
	}

	//Send alert email
	sendEmailAlert(containers)
}
