package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/iwaltgen/sak/cmd/sak/rand"
)

func main() {
	app := &cli.App{
		Name:  "sak",
		Usage: "multi-tool cli",
		Commands: []*cli.Command{
			rand.Cmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
