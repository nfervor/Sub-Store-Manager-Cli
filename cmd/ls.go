package cmd

import (
	"fmt"

	"sub-store-manager-cli/docker"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all sub-store docker containers",
	Run: func(cmd *cobra.Command, args []string) {
		listAllSSMContainer()
	},
}

func listAllSSMContainer() {
	fel, bel := docker.GetSSMContainers()

	if len(fel) == 0 && len(bel) == 0 {
		fmt.Println("No Sub-Store Manager Docker Containers found")
		return
	}

	fmt.Println("Sub-Store Manager Docker Containers:")
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Type", "ID", "Version", "Ports", "Status", "Names")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, c := range append(fel, bel...) {
		var portStr string
		if p, e := c.GetPortInfo(); e != nil {
			portStr = "none"
		} else {
			portStr = fmt.Sprintf("%s:%s->%s/%s", p.HostIP, p.Public, p.Private, p.Type)
		}
		tbl.AddRow(c.ContainerType, c.DockerContainer.ID[:24], c.Version, portStr, c.DockerContainer.State, c.Name)
	}

	tbl.Print()
}
