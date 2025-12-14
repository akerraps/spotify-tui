package main

import (
	"context"
	"fmt"
	"log"
	"os"

	Core "akerraps/tunectl/internal/cli"

	"github.com/joho/godotenv"
	urfave "github.com/urfave/cli/v2"
)

func init() {
	godotenv.Load()
}

func NewApp(ctx context.Context) *Core.Core {
	return &Core.Core{
		Name: "TuneCtl",
	}
}

func main() {
	if len(os.Args) > 1 {
		runTui()
	}
}

func runTui() {
	ctx := context.Background()

	appCore := NewApp(ctx)

	cliApp := &urfave.App{
		Name: "TuneCtl",
		Commands: []*urfave.Command{
			{
				Name: "playlists",
				Action: func(c *urfave.Context) error {
					return appCore.RunPlaylists(c.Context)
				},
			},
			{
				Name:      "songs",
				Usage:     "Lists songs from a specefied playlist",
				ArgsUsage: "<playlist>",
				Action: func(c *urfave.Context) error {
					if c.NArg() < 1 {
						return fmt.Errorf("you must specify a playlist")
					}

					playlist := c.Args().Get(0)

					return appCore.RunSogns(c.Context, playlist)
				},
			},
		},
	}

	err := cliApp.RunContext(ctx, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
