package main

import (
	"fmt"
	"os"

	"github.com/gridprotocol/platform-v2/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	local := make([]*cli.Command, 0, 1)
	local = append(local, cmd.DaemonCmd)
	local = append(local, cmd.VersionCmd)
	local = append(local, cmd.TopupCmd)
	local = append(local, cmd.Topup2Cmd)

	app := cli.App{
		Commands: local,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Show application version",
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Bool("version") {
				fmt.Println(cmd.Version + "+" + cmd.BuildFlag)
			}
			return nil
		},
	}
	app.Setup()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err) // nolint:errcheck
		os.Exit(1)
	}
}
