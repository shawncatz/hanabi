package main

import "gopkg.in/urfave/cli.v1"

func init() {
	cmdList = append(cmdList, cli.Command{
		Name:  "service",
		Usage: "manage and create services",
		Description: "manage and create services",
		Action:    cmdServiceCreate,
		ArgsUsage: "",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "create",
				Action: cmdServiceCreate,
			},
		},
	})
}

func cmdServiceCreate(c *cli.Context) error {
	return nil
}