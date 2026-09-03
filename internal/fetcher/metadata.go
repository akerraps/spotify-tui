package fetcher

import (
	"log/slog"

	"akerraps/tunectl/internal/types"

	"go.senan.xyz/taglib"
)

func writeMetadata(file string, song *types.TrackInfo) {
	log := slog.With(
		"track_id", song.ID,
		"title", song.Title,
		"artists", song.Artists,
	)

	tags := map[string][]string{
		taglib.Album:              {song.Album},
		taglib.Artist:             song.Artists,
		taglib.Genre:              song.Genres,
		taglib.Title:              {song.Title},
		taglib.MusicBrainzTrackID: {song.MBID},
	}

	err := taglib.WriteTags(file, tags, 0)
	if err != nil {
		log.Warn(
			"failed to write metadata",
			"file", file,
			"err", err,
		)

		return
	}

	log.Debug(
		"metadata written",
		"genres", song.Genres,
		"album", song.Album,
	)
}
