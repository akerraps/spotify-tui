package fetcher

import (
	"fmt"
	"log/slog"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"

	"akerraps/tunectl/internal/musicbrainz"
	"akerraps/tunectl/internal/types"

	"golang.org/x/text/unicode/norm"
)

func outputPath(song types.TrackInfo, outputDir string) string {
	name := song.Title + "_" + song.Artists[0]

	name = norm.NFD.String(name)

	var b strings.Builder

	for _, r := range name {
		if unicode.Is(unicode.Mn, r) {
			continue
		}

		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' || r == '_' {
			b.WriteRune(r)
		}
	}

	name = strings.ToLower(b.String())
	name = strings.ReplaceAll(name, " ", "_")

	return filepath.Join(outputDir, name+".mp3")
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

		if err := musicbrainz.GetSongInfo(&song); err != nil {
			slog.Warn(
				"failed to fetch metadata, using original data",
				"title", song.Title,
				"artists", song.Artists,
				"err", err,
			)
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
