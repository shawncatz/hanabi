package main

import "gopkg.in/urfave/cli.v1"

func init() {
	cmdList = append(cmdList, cli.Command{
		Name:  "config",
		Usage: "manage configuration",
		Description: "manage configuration",
		Action:    cmdConfigList,
		ArgsUsage: "",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "list",
				Action: cmdConfigList,
			},
		},
	})
}

func cmdConfigList(c *cli.Context) error {
	return nil
}
