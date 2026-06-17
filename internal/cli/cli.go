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
						parts := strings.SplitN(song, ";", 3)

						track := types.TrackInfo{}

						if len(parts) > 0 {
							track.Title = strings.TrimSpace(parts[0])
						}

						if len(parts) > 1 {
							artists := []string{strings.TrimSpace(parts[1])}
							for i := range artists {
								artists[i] = strings.TrimSpace(artists[i])
							}
							track.Artists = artists
						}

						if len(parts) > 2 {
							track.Genres = []string{strings.TrimSpace(parts[2])}
						}

						tracks = append(tracks, track)
					}

					return fetcher.FetchAudio(tracks, opts)

				},
			},
			{
				Name:  "file",
				Usage: "Download data from a file",
				Flags: []urfave.Flag{
					&urfave.BoolFlag{
						Name:  "csv",
						Usage: "Download from CSV",
					},
					&urfave.BoolFlag{
						Name:  "json",
						Usage: "Download from JSON",
					},
					&urfave.StringFlag{
						Name:    "data",
						Aliases: []string{"d"},
						Usage:   "Data file path",
					}, &urfave.StringFlag{
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

					csv := c.Bool("csv")
					json := c.Bool("json")

					if csv == json {
						return fmt.Errorf("You must choose file type")
					}

					file := c.String("data")

					var tracks []types.TrackInfo
					var err error

					if csv == true {
						tracks, err = fetcher.ReadCsvFile(file)
					} else {
						tracks, err = fetcher.ReadJsonFile(file)
					}

					if err != nil {
						log.Fatal(err)
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
