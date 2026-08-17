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

func parallelFlag() urfave.Flag {
	return &urfave.IntFlag{
		Name:    "paralelism",
		Aliases: []string{"n"},
		Usage:   "Amount of parallel downloads",
	}
}

func buildOptions(c *urfave.Context) types.Options {
	opts := types.DefaultOptions()

	if out := c.String("output"); out != "" {
		opts.OutputDir = out

		slog.Debug(
			"using custom output directory",
			"output_dir", opts.OutputDir,
		)
	} else {
		slog.Debug(
			"using default output directory",
			"output_dir", opts.OutputDir,
		)
	}

	if c.IsSet("paralelism") {
		opts.Parallelism = c.Int("paralelism")

		slog.Debug(
			"using custom parallelism",
			"parallelism", opts.Parallelism,
		)
	} else {
		slog.Debug(
			"using default parallelism",
			"parallelism", opts.Parallelism,
		)
	}

	opts.NoAPI = c.Bool("no-api")

	if opts.NoAPI {
		slog.Debug("musicbrainz metadata lookup disabled")
	}

	return opts
}
