package cli

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"akerraps/tunectl/internal/cache"
	"akerraps/tunectl/internal/core"
	"akerraps/tunectl/internal/fetcher"
	"akerraps/tunectl/internal/types"

	urfave "github.com/urfave/cli/v2"
)

func NewApp(ctx context.Context) *core.Service {
	return &core.Service{
		Name: "TuneCtl",
	}
}

func RunCli() {
	ctx := context.Background()

	appCore := NewApp(ctx)

	cliApp := &urfave.App{
		Name:  "tunectl",
		Usage: "Manage your playlists and songs",
		Commands: []*urfave.Command{
			{
				Name:  "cache",
				Usage: "Manage cache",
				Flags: []urfave.Flag{
					&urfave.BoolFlag{
						Name:    "clear",
						Aliases: []string{"c"},
						Usage:   "Clear cache",
					},
				},
				Action: func(c *urfave.Context) error {
					if c.Bool("clear") {
						return cache.ClearYtDlp()
					} else {
						return fmt.Errorf("no valid flag provided")
					}
				},
			},
			{
				Name:  "playlists",
				Usage: "Manage playlists",
				Flags: []urfave.Flag{
					&urfave.BoolFlag{
						Name:    "list",
						Aliases: []string{"l"},
						Usage:   "List all playlists",
					},
					&urfave.BoolFlag{
						Name:    "download",
						Aliases: []string{"d"},
						Usage:   "Download songs from a playlist",
					},
					&urfave.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "Directory where songs will be downloaded",
					},
				},
				Action: func(c *urfave.Context) error {
					if c.Bool("list") {
						return appCore.RunPlaylists(c.Context)
					}

					if c.Bool("download") {
						if c.NArg() < 1 {
							return fmt.Errorf("you must specify a playlist")
						}
						playlist := c.Args().Get(0)

						out := c.String("output")
						if out == "" {
							return fmt.Errorf("the output directory must be specified")
						}
						return appCore.RunSongs(c.Context, playlist, true, out)
					}

					return fmt.Errorf("no valid flag provided")
				},
			},
			{
				Name:      "songs",
				Usage:     "Manage songs",
				ArgsUsage: "<playlist>",
				Flags: []urfave.Flag{
					&urfave.BoolFlag{
						Name:    "list",
						Aliases: []string{"l"},
						Usage:   "List songs in a playlist",
					},
					&urfave.BoolFlag{
						Name:    "download",
						Aliases: []string{"d"},
						Usage:   "Download a song by name",
					},
					&urfave.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "Directory where songs will be downloaded",
					},
				},
				Action: func(c *urfave.Context) error {
					if c.NArg() < 1 {
						return fmt.Errorf("you must specify a playlist")
					}

					out := c.String("output")

					if c.Bool("download") {
						if out == "" {
							return fmt.Errorf("the output directory must be specified")
						}

						if c.NArg() == 0 {
							return fmt.Errorf("you must specify at least one song")
						}

						args := c.Args().Slice()

						tracks := make([]types.TrackInfo, 0, len(args))
						for _, song := range args {
							name := strings.Split(song, ";")[0]
							artist := []string{}
							artist = append(artist, strings.Split(song, ";")[1])
							tracks = append(tracks, types.TrackInfo{
								Name:    name,
								Artists: artist,
							})
						}

						return fetcher.FetchAudio(tracks, out)

					} else if c.Bool("list") {
						playlist := c.Args().Get(0)
						return appCore.RunSongs(c.Context, playlist, false, out)
					}

					return fmt.Errorf("no valid flag provided")
				},
			},
		},
	}

	err := cliApp.RunContext(ctx, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
