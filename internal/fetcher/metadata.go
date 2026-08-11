package fetcher

import (
	"akerraps/tunectl/internal/types"
	"log/slog"
	"strconv"

	"go.senan.xyz/taglib"
)

func writeMetadata(file string, song types.TrackInfo) {
	tags := map[string][]string{
		taglib.Album:  {song.Album},
		taglib.Artist: song.Artists,
		taglib.Genre:  song.Genres,
		taglib.Title:  {song.Title},
	}

	if song.Year != 0 {
		tags[taglib.Date] = []string{strconv.Itoa(song.Year)}
	}

	err := taglib.WriteTags(file, tags, 0)
	if err != nil {
		slog.Warn(
			"failed to write metadata",
			"title", song.Title,
			"artists", song.Artists,
			"file", file,
			"err", err,
		)

		return
	}

	slog.Debug(
		"metadata written",
		"title", song.Title,
		"artists", song.Artists,
		"genres", song.Genres,
		"year", song.Year,
		"album", song.Album,
	)
}
