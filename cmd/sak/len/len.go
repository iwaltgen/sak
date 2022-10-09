// Copyright (c) 2022 iwaltgen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package len

import (
	"bufio"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/urfave/cli/v2"
)

// Cmd represents the task command
var Cmd *cli.Command

func init() {
	Cmd = &cli.Command{
		Name:   "len",
		Usage:  "calculate a string length",
		Action: run,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "base64",
				Value: false,
				Usage: "decoding base64",
			},
			&cli.BoolFlag{
				Name:  "base32",
				Value: false,
				Usage: "decoding base32",
			},
			&cli.BoolFlag{
				Name:  "hex",
				Value: false,
				Usage: "decoding hex",
			},
		},
	}
}

func run(ctx *cli.Context) error {
	text := ctx.Args().First()
	if text == "" {
		scanner := bufio.NewScanner(ctx.App.Reader)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("read line error: %w", err)
		}

		text = scanner.Text()
	}

	switch {
	case ctx.Bool("hex"):
		bytes, err := hex.DecodeString(text)
		if err != nil {
			return fmt.Errorf("decode hex error: %w", err)
		}
		text = string(bytes)

	case ctx.Bool("base32"):
		bytes, err := base32.StdEncoding.DecodeString(text)
		if err != nil {
			return fmt.Errorf("decode base32 error: %w", err)
		}
		text = string(bytes)

	case ctx.Bool("base64"):
		bytes, err := base64.URLEncoding.DecodeString(text)
		if err != nil {
			return fmt.Errorf("decode base64 error: %w", err)
		}
		text = string(bytes)
	}

	fmt.Fprint(ctx.App.Writer, len(text))
	fmt.Fprintln(ctx.App.ErrWriter)
	return nil
}
