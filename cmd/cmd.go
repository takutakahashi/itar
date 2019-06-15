package main

import (
	"github.com/urfave/cli"
	"os"
  "github.com/docker/docker/pkg/archive"
  "github.com/docker/docker/builder/dockerignore"
  "io"
  "fmt"
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
    fmt.Println("start app")
		return makeTar(".", c.Args().Get(0))
	}
	return app
}

func makeTar(dir string, ignore string) error {
  f, err := os.Open(ignore)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  ignorePatterns, err := dockerignore.ReadAll(f)
  tarStream, err := archive.TarWithOptions(dir, &archive.TarOptions{
    ExcludePatterns: ignorePatterns,
    Compression: archive.Uncompressed})
  if err != nil {
    panic(err)
  }
  defer tarStream.Close()
  outputFile, err := os.Create("context")
  defer outputFile.Close()
  io.Copy(outputFile, tarStream)
	return nil
}

