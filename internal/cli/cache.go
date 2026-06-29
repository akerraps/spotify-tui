package cli

import (
	"akerraps/tunectl/internal/cache"
	"fmt"

	urfave "github.com/urfave/cli/v2"
)

func cacheCommand() *urfave.Command {
	return &urfave.Command{
		Name:  "cache",
		Usage: "Manage cache",

		Flags: []urfave.Flag{
			&urfave.BoolFlag{
				Name:    "clear",
				Aliases: []string{"c"},
				Usage:   "Clear cache",
			},
		},

		Action: cacheAction,
	}
}

func cacheAction(c *urfave.Context) error {
	if c.Bool("clear") {
		return cache.ClearYtDlp()
	}

	return fmt.Errorf("no valid flag provided")
}
