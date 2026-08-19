package cli

import (
	"akerraps/tunectl/internal/types"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	urfave "github.com/urfave/cli/v2"
)

func outputFlag() urfave.Flag {
	return &urfave.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   "Directory where songs will be downloaded",
	}
}

func envFlag() urfave.Flag {
	return &urfave.StringFlag{
		Name:    "environment",
		Aliases: []string{"e"},
		Usage:   "Path to the environment file",
		Value:   ".env",
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
		Value:   1,
	}
}

func buildOptions(c *urfave.Context) types.Options {
	opts := types.DefaultOptions()

	if envFile := c.String("environment"); envFile != "" {
		opts.EnvFile = envFile

		slog.Debug(
			"overriding environment file with CLI option",
			"env_file", opts.EnvFile,
		)
	}

	readEnv(&opts)

	if out := c.String("output"); out != "" {
		opts.OutputDir = out

		slog.Debug(
			"overriding output directory with CLI option",
			"output_dir", opts.OutputDir,
		)
	}

	if c.IsSet("paralelism") {
		opts.Parallelism = c.Int("paralelism")

		slog.Debug(
			"overriding parallelism with CLI option",
			"parallelism", opts.Parallelism,
		)
	}

	if c.IsSet("no-api") {
		opts.NoAPI = c.Bool("no-api")

		slog.Debug(
			"overriding no-api with CLI option",
			"no_api", opts.NoAPI,
		)
	}

	slog.Debug(
		"final options",
		"output_dir", opts.OutputDir,
		"no_api", opts.NoAPI,
		"parallelism", opts.Parallelism,
	)

	return opts
}

func readEnv(opts *types.Options) {
	if err := godotenv.Overload(opts.EnvFile); err != nil {
		slog.Debug(
			"failed to load environment file",
			"file", opts.EnvFile,
			"error", err,
		)
		return
	}

	if err := envconfig.Process("TUNECTL", opts); err != nil {
		slog.Debug(
			"failed to process environment configuration",
			"error", err,
		)
	}
}
