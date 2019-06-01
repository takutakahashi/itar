package main

import (
	"github.com/urfave/cli"
	"os"
)

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
	app.Action = func(c *cli.Context) error {
		return nil
	}
	return app
}
