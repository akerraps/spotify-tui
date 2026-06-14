package cli

import (
	"akerraps/tunectl/internal/cache"
	"akerraps/tunectl/internal/fetcher"
	"akerraps/tunectl/internal/types"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
				Name:      "songs",
				Usage:     "Download songs",
				ArgsUsage: "<song>",
				Flags: []urfave.Flag{

					&urfave.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "Directory where songs will be downloaded",
					},
					&urfave.BoolFlag{
						Name:  "no-api",
						Usage: "Omit MusicBrainz api usage",
					},
				},

				Action: func(c *urfave.Context) error {

					if c.NArg() == 0 {
						return fmt.Errorf("you must specify at least one song")
					}

					opts := types.DefaultOptions()

					if out := c.String("output"); out != "" {
						opts.OutputDir = out
					} else {
						log.Printf(
							"warning: output directory not specified, using default: %s",
							opts.OutputDir,
						)
					}

					opts.NoAPI = c.Bool("no-api")

					args := c.Args().Slice()

					tracks := make([]types.TrackInfo, 0, len(args))
					for _, song := range args {

						name, artistName, found := strings.Cut(song, ";")

						artists := []string{}
						if found {
							artists = append(artists, artistName)
						}

						tracks = append(tracks, types.TrackInfo{
							Title:   name,
							Artists: artists,
						})
					}

					return fetcher.FetchAudio(tracks, opts)

				},
			},
		},
	}

	err := cliApp.RunContext(ctx, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
