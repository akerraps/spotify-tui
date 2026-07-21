package fetcher

import (
	"log/slog"
	"path/filepath"
	"strings"

	"akerraps/tunectl/internal/cache"
	"akerraps/tunectl/internal/types"
)

func songExists(path string) (bool, error) {
	prefix := strings.TrimSuffix(path, filepath.Ext(path))

	matches, err := filepath.Glob(prefix + ".*")
	if err != nil {
		return false, err
	}

	return len(matches) > 0, nil
}

func FetchAudio(tracks []types.TrackInfo, opts types.Options) error {
	bin, err := cache.GetYtDlp()

	if err != nil {
		return err
	}

	slog.Debug(
		"using yt-dlp binary",
		"path", bin,
	)

	for _, song := range tracks {
		if err := processTrack(bin, song, opts); err != nil {
			slog.Error(
				"failed to process track",
				"title", song.Title,
				"err", err,
			)
		}
	}
	return nil
}
