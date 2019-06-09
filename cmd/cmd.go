package main

import (
	"fmt"
	"github.com/mholt/archiver"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path/filepath"
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
		return makeTar(c.Args().Get(0))
	}
	return app
}

func makeTar(archivePath string) error {
	tarPath := archivePath + ".tar"
	tar := archiver.NewTar()
	filePaths := getFileDirList(archivePath)
	fmt.Println(filePaths)
	tar.Archive(filePaths, tarPath)
	return nil
}

func getFileDirList(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, getFileDirList(filepath.Join(path, file.Name()))...)
			continue
		} else {
			paths = append(paths, filepath.Join(path, file.Name()))
		}
	}
	return paths
}
