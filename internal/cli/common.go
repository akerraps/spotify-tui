package cli

import (
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
