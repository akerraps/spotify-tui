package cli

import (
	"akerraps/tunectl/internal/types"
	"log/slog"

	urfave "github.com/urfave/cli/v2"
)

func outputFlag() urfave.Flag {
	return &urfave.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   "Directory where songs will be downloaded",
	}
}

func noAPIFlag() urfave.Flag {
	return &urfave.BoolFlag{
		Name:  "no-api",
		Usage: "Omit MusicBrainz API usage",
	}
}

func buildOptions(c *urfave.Context) types.Options {
	opts := types.DefaultOptions()

	if out := c.String("output"); out != "" {
		opts.OutputDir = out
	} else {
		slog.Warn(
			"output directory not specified, using default",
			"output_dir", opts.OutputDir,
		)
	}

	opts.NoAPI = c.Bool("no-api")

	return opts
}
