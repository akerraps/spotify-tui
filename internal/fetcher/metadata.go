package fetcher

import (
	"akerraps/tunectl/internal/types"
	"log/slog"

	"go.senan.xyz/taglib"
)

func writeMetadata(file string, song types.TrackInfo) {

	err := taglib.WriteTags(file, map[string][]string{
		taglib.Album:  {song.Album},
		taglib.Artist: song.Artists,
		taglib.Genre:  song.Genres,
		taglib.Title:  {song.Title},
	}, 0)
	if err != nil {
		slog.Warn(
			"failed to write metadata",
			"title", song.Title,
			"artists", song.Artists,
			"file", file,
			"err", err,
		)
	} else {
		slog.Debug(
			"metadata writen",
			"title", song.Title,
			"artists", song.Artists,
			"genres", song.Genres,
			"album", song.Album,
		)
	}
}
