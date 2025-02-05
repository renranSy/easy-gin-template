package main

import (
	"easy-gin-template/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

var VERSION = "v1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "easy-gin-template"
	app.Version = VERSION
	app.Usage = "Sever Start"
	app.Commands = []*cli.Command{
		cmd.StartCmd(),
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
