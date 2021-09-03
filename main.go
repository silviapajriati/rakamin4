package main

import (
	"os"
	"rakamin4/app"
	"rakamin4/cli"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
