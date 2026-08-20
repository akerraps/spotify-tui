package cli

import (
	"akerraps/tunectl/internal/fetcher"
	"akerraps/tunectl/internal/types"
	"fmt"
	"strings"

	urfave "github.com/urfave/cli/v2"
)

func songsCommand() *urfave.Command {
	return &urfave.Command{
		Name:      "songs",
		Usage:     "Download songs",
		ArgsUsage: "<song>",

		Flags: []urfave.Flag{
			outputFlag(),
			envFlag(),
			parallelFlag(),
			rewriteMetadataFlag(),
			noAPIFlag(),
		},

		Action: songsAction,
	}
}

func songsAction(c *urfave.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("you must specify at least one song")
	}

	opts := buildOptions(c)

	tracks, err := parseSongs(c.Args().Slice())
	if err != nil {
		return err
	}

	return fetcher.FetchAudio(tracks, opts)
}

func parseSongs(args []string) ([]types.TrackInfo, error) {
	tracks := make([]types.TrackInfo, 0, len(args))

	for _, song := range args {
		parts := strings.SplitN(song, ";", 3)

		track := types.TrackInfo{}

		if len(parts) >= 1 {
			track.Title = strings.TrimSpace(parts[0])
		}

		if len(parts) >= 2 {
			track.Artists = splitAndTrim(parts[1])
		}

		if len(parts) >= 3 {
			track.Genres = splitAndTrim(parts[2])
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func splitAndTrim(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}

	parts := strings.Split(s, ",")

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return parts
}
