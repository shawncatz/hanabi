package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

var (
	cmdList                   = []cli.Command{} // so that the individual files can append to the list
)

func main() {
	args := os.Args

	app := cli.NewApp()
	app.Name = "hanabi"
	app.Version = "0.1.0"
	app.Author = "Shawn Catanzarite"
	app.Usage = "CLI for hanabi server"
	app.UsageText = "hanabi [global options] <command> [command options] [arguments...]"
	app.Email = "me@shawncatz.com"
	app.Description = "interact with the hanabi server, create new services, manage configuration"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "no-color",
			Usage: "disable color output",
		},
	}
	app.Commands = cmdList

	app.Run(args)
}