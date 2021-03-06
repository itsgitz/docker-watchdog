package watchdog

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
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

func (w *Watchdog) getStoppedContainers() {
	ctx := context.Background()

	//Create new API client
	//WithAPIVersionNegotiation enables automatic API version negotiation for the client
	//With this option enabled, the client automatically negotiates the API version to use when making requests
	//src: https://pkg.go.dev/github.com/docker/docker/client#WithAPIVersionNegotiation
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		cobra.CheckErr(err)
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
		cobra.CheckErr(err)
	}

	InformationText.Printf("[*] Searching for stopped containers (exited|paused|dead) ... \n")

	//Return nothing if there are not stopped containers
	//And continue detection ...
	if len(containers) == 0 {
		SuccessText.Printf("[*] All containers are running ... \n\n")
		return
	}

	//Display all stopped containers information
	DangerText.Printf("\n[!] Stopped container is detected! \n\n")
	for _, c := range containers {
		DangerText.Printf("| ID: %v \n", c.ID[:10])
		DangerText.Printf("| Name: %v \n", strings.Trim(c.Names[0], "/"))
		DangerText.Printf("| Status: %v \n", c.Status)
		DangerText.Printf("| State: %v \n\n", c.State)
	}

	//Send alert email
	w.sendEmailAlert(containers)
}
