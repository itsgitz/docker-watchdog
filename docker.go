package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

//Define not running containers state
//status=(paused|exited|dead)
//src: https://docs.docker.com/engine/api/v1.41/#operation/ContainerList
const (
	ExitedState = "exited" //Exited
	PausedState = "paused" //Paused
	DeadState   = "dead"   //Dead
)

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

	fmt.Printf("Search for not running containers (exited|paused|dead) ... \n")

	//
	if len(containers) == 0 {
		fmt.Println("All containers are running ...")
		return
	}

	fmt.Printf("\nFound not running containers! \n\n")
	for _, c := range containers {
		fmt.Printf("ID: %v \n", c.ID[:10])
		fmt.Printf("Name: %v \n", strings.Trim(c.Names[0], "/"))
		fmt.Printf("Status: %v \n", c.Status)
		fmt.Printf("State: %v \n\n", c.State)
	}
}
