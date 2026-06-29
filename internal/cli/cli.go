package cli

import (
	"context"
	"log/slog"
	"os"

	urfave "github.com/urfave/cli/v2"
)

type Service struct {
	Name string
}

func NewApp(ctx context.Context) *Service {
	return &Service{
		Name: "TuneCtl",
	}
}

func RunCli() {
	ctx := context.Background()

	app := &urfave.App{
		Name:  "tunectl",
		Usage: "Manage your playlists and songs",

		Flags: []urfave.Flag{
			&urfave.BoolFlag{
				Name:  "debug",
				Usage: "enable debug logs",
			},
		},

		Before: setupLogger,

		Commands: []*urfave.Command{
			cacheCommand(),
			songsCommand(),
			fileCommand(),
		},
	}

	if err := app.RunContext(ctx, os.Args); err != nil {
		slog.Error(
			"command execution failed",
			"err", err,
		)
		os.Exit(1)
	}
}

func setupLogger(c *urfave.Context) error {
	level := slog.LevelInfo
	if c.Bool("debug") {
		level = slog.LevelDebug
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	slog.SetDefault(slog.New(handler))

	return nil
}
