package fetcher

import (
	"fmt"
	"log/slog"
	"os/exec"
	"path/filepath"
	"strings"

	"akerraps/tunectl/internal/musicbrainz"
	"akerraps/tunectl/internal/types"
)

func outputPath(song types.TrackInfo, OutputDir string) string {
	fileName := strings.ToLower(
		strings.ReplaceAll(
			song.Title+"_"+song.Artists[0]+".mp3",
			" ",
			"_",
		),
	)

	return filepath.Join(OutputDir, fileName)
}

func processTrack(bin string, song types.TrackInfo, opts types.Options) error {
	slog.Debug(
		"processing track",
		"title", song.Title,
		"artists", song.Artists,
	)

	output := outputPath(song, opts.OutputDir)

	exists, err := songExists(output)
	if err != nil {
		return err
	}

	if !opts.NoAPI {
		slog.Debug(
			"fetching metadata from api",
			"title", song.Title,
			"artists", song.Artists,
		)

		info, err := musicbrainz.GetSongInfo(song)
		if err != nil {
			slog.Warn(
				"failed to fetch metadata, using original data",
				"title", song.Title,
				"artists", song.Artists,
				"err", err,
			)
		} else {
			song.Title = info.Title
			song.Artists = info.Artists
			song.Genres = append(song.Genres, info.Genres...)
		}
	}

	if exists {
		slog.Info(
			"song already exists",
			"title", song.Title,
			"artists", song.Artists,
		)

		writeMetadata(output, song)

		return nil
	}

	search := fmt.Sprintf(
		`ytsearch:"%s" %s "song"`,
		song.Title,
		strings.Join(song.Artists, " "),
	)

	slog.Debug(
		"starting download",
		"title", song.Title,
		"artists", song.Artists,
		"query", search,
	)

	cmd := exec.Command(
		bin,
		"-x",
		"--restrict-filenames",
		"--quiet",
		"--no-warnings",
		"--embed-thumbnail",
		"--no-playlist",
		"-t", "mp3",
		"--audio-quality", "128K",
		search,
		"-o", output,
	)

	if _, err := cmd.Output(); err != nil {
		slog.Error(
			"download failed",
			"title", song.Title,
			"artists", song.Artists,
			"err", err,
		)
		return nil
	}

	slog.Info(
		"download completed",
		"title", song.Title,
		"artists", song.Artists,
	)

	writeMetadata(output, song)

	return nil
}
