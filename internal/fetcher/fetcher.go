package fetcher

import (
	"log/slog"
	"path/filepath"
	"strings"
	"time"

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

	total := len(tracks)
	start := time.Now()

	for i, song := range tracks {
		current := i + 1
		remaining := total - current

		if err := processTrack(bin, song, opts); err != nil {
			slog.Error(
				"failed to process track",
				"title", song.Title,
				"err", err,
			)
		}

		elapsed := time.Since(start)
		avgPerTrack := elapsed / time.Duration(current)
		eta := avgPerTrack * time.Duration(remaining)

		slog.Info(
			"track progress",
			"current", current,
			"total", total,
			"remaining", remaining,
			"elapsed", elapsed.Round(time.Second),
			"eta", eta.Round(time.Second),
		)

		slog.Info("") // For reading log easily per song

	}

	return nil
}
