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
		Usage:  "create a new UUID string (16 bytes, 36 runes)",
		Action: run,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "xid",
				Value: false,
				Usage: "encoding XID (12 bytes, 20 runes)",
			},
			&cli.BoolFlag{
				Name:  "base64",
				Value: false,
				Usage: "encoding base64 (16 bytes, 22 runes)",
			},
			&cli.BoolFlag{
				Name:  "base62",
				Value: false,
				Usage: "encoding base62 (16 bytes, 22 runes)",
			},
			&cli.BoolFlag{
				Name:  "ulid",
				Value: false,
				Usage: "engine use ULID (16 bytes, 26 runes)",
			},
			&cli.BoolFlag{
				Name:  "ksuid",
				Value: false,
				Usage: "engine use KSUID (20 bytes, 27 runes)",
			},
			&cli.BoolFlag{
				Name:  "hex",
				Value: false,
				Usage: "encoding hex (16 bytes, 32 runes)",
			},
		},
	}
}

func run(ctx *cli.Context) error {
	var out string
	switch {
	case ctx.Bool("xid"):
		out = uuid.NewXID()

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
