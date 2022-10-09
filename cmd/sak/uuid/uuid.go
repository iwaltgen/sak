// Copyright (c) 2022 iwaltgen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package uuid

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/iwaltgen/sak/pkg/uuid"
)

// Cmd represents the task command
var Cmd *cli.Command

func init() {
	Cmd = &cli.Command{
		Name:   "uuid",
		Usage:  "create a new UUID string",
		Action: run,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "base64",
				Value: false,
				Usage: "encoding base64 (16bytes, 22runes)",
			},
			&cli.BoolFlag{
				Name:  "base62",
				Value: false,
				Usage: "encoding base62 (16bytes, 22runes)",
			},
			&cli.BoolFlag{
				Name:  "ulid",
				Value: false,
				Usage: "engine use ULID (16bytes, 26runes)",
			},
			&cli.BoolFlag{
				Name:  "ksuid",
				Value: false,
				Usage: "engine use KSUID (20bytes, 27runes)",
			},
			&cli.BoolFlag{
				Name:  "hex",
				Value: false,
				Usage: "encoding hex (16bytes, 32runes)",
			},
		},
	}
}

func run(ctx *cli.Context) error {
	var out string
	switch {
	case ctx.Bool("base64"):
		out = uuid.NewBase64()

	case ctx.Bool("base62"):
		out = uuid.NewBase62()

	case ctx.Bool("ulid"):
		out = uuid.NewULID()

	case ctx.Bool("ksuid"):
		out = uuid.NewKSUID()

	case ctx.Bool("hex"):
		out = uuid.NewHex()

	default:
		out = uuid.New()
	}

	fmt.Fprint(ctx.App.Writer, out)
	fmt.Fprintln(ctx.App.ErrWriter)
	return nil
}
