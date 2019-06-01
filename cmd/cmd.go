package main

import (
	"os"
)

var Commands = []cli.Command{
	commandClient,
}

var commandClient = cli.Command{
	Name:        "client",
	Usage:       "usage",
	Description: "desc",
	Action:      doClient,
}

func doClient(c *cli.Context) {
	client.ConnectServer()
}

var Version string = "0.9.0"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "itar"
	app.Usage = "ignorable tar archiver"
	app.Version = Version
	app.Author = "takutakahashi"
	app.Email = "taku.takahashi120@gmail.com"
	app.Commands = Commands
	return app
}
