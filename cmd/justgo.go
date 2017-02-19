package main // import "github.com/inadarei/justgo/cmd"

import (
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "justgo"
	app.Usage = "create a new skeleton project for a Go-based API [micro]service"
	app.Action = func(c *cli.Context) error {
		installPath := "."
		if c.NArg() > 0 {
			installPath = c.Args().Get(0)
		}
		buildProject(installPath)
		return nil
	}

	app.Run(os.Args)

}

func buildProject(path string) {
	fmt.Println("Building a skeleton microservice at:", path)
	exists, err := pathExists(path)
	abortIfErr(err)

	if exists == true { // need to make sure it's empty!
		isPathEmpty := folderIsEmpty(path)
		if !isPathEmpty {
			fmt.Println("There are files in ", path, " Destination path must be empty. Aborting.")
		}
	} else {
		err := os.MkdirAll(path, os.ModePerm)
		abortIfErr(err)
	}
}

func abortIfErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func folderIsEmpty(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		f.Close()
		abortIfErr(err)
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true
	}
	// If not already exited, scanning children must have errored-out
	abortIfErr(err)
	return false
}
