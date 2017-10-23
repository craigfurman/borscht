package main

import (
	"fmt"
	"os"

	"github.com/craigfurman/borscht/borscht"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "borscht"
	app.Usage = "See the diff of bosh jobs between releases"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "from,f", Usage: "Old bosh final release version"},
		cli.StringFlag{Name: "to,t", Usage: "New bosh final release version"},
	}

	app.Action = run
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}

func run(ctx *cli.Context) error {
	fromVersion, err := mustHaveFlag(ctx, "from")
	if err != nil {
		return err
	}
	toVersion, err := mustHaveFlag(ctx, "to")
	if err != nil {
		return err
	}
	if ctx.NArg() != 1 {
		return fmt.Errorf("expected 1 non-flag argument, got %d", ctx.NArg())
	}
	releasePath := ctx.Args().Get(0)

	jobDiffs, err := borscht.Diff(releasePath, fromVersion, toVersion)
	if err != nil {
		return err
	}

	for job, diff := range jobDiffs {
		fmt.Printf("%s:\n%s\n", job, diff)
	}

	return nil
}

func mustHaveFlag(ctx *cli.Context, key string) (string, error) {
	if !ctx.IsSet(key) {
		return "", fmt.Errorf("must pass the '--%s' flag", key)
	}
	return ctx.String(key), nil
}
