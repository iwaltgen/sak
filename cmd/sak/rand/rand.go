package rand

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/iwaltgen/sak/pkg/rand"
)

// Cmd represents the task command
var Cmd *cli.Command

func init() {
	Cmd = &cli.Command{
		Name:   "rand",
		Usage:  "create a new random string",
		Action: run,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "length",
				Aliases: []string{"l"},
				Value:   16,
				Usage:   "length of random bytes",
			},
			&cli.BoolFlag{
				Name:  "base64",
				Value: true,
				Usage: "encoding base64",
			},
			&cli.BoolFlag{
				Name:  "base32",
				Value: false,
				Usage: "encoding base32",
			},
			&cli.BoolFlag{
				Name:  "hex",
				Value: false,
				Usage: "encoding hex",
			},
			&cli.BoolFlag{
				Name:  "ksuid",
				Value: false,
				Usage: "engine use KSUID (fixed length: 20byte)",
			},
			&cli.BoolFlag{
				Name:  "ulid",
				Value: false,
				Usage: "engine use ULID (fixed length: 16byte)",
			},
		},
	}
}

func run(ctx *cli.Context) error {
	var (
		random string
		length = ctx.Int("length")
	)

	switch {
	case ctx.Bool("hex"):
		random = rand.NewHexWithLength(length)

	case ctx.Bool("base32"):
		random = rand.NewBase32WithLength(length)

	case ctx.Bool("ksuid"):
		random = rand.NewKSUID()

	case ctx.Bool("ulid"):
		random = rand.NewULID()

	default:
		random = rand.NewBase64WithLength(length)
	}

	fmt.Fprint(ctx.App.Writer, random)
	fmt.Fprintln(ctx.App.ErrWriter)
	return nil
}
