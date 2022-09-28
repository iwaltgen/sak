package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/iwaltgen/sak/cmd/sak/rand"
	"github.com/iwaltgen/sak/internal"
)

func main() {
	app := &cli.App{
		Name:     "sak",
		Usage:    "Multi-tool CLI",
		Version:  internal.Version(),
		Compiled: internal.BuildDate(),
		Authors: []*cli.Author{
			{
				Name:  "iwaltgen",
				Email: "iwaltgen@gmail.com",
			},
		},

		Suggest:                true,
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,

		Before: func(ctx *cli.Context) error {
			log.SetFlags(0)
			return nil
		},
		Commands: []*cli.Command{
			rand.Cmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
