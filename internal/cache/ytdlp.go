package cache

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-getter"
)

func GetYtDlp() (string, error) {
	binary := "yt-dlp"
	cachePath, err := getCachePath()
	if err != nil {
		return "", fmt.Errorf("cannot get cache path: %w", err)
	}
	path := filepath.Join(cachePath, binary)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		slog.Info(
			"yt-dlp not found in cache, downloading",
			"path", path,
		)

		slog.Debug(
			"ensuring cache directory exists",
			"dir", cachePath,
		)

		if err := createCacheDir(cachePath); err != nil {
			return "", fmt.Errorf("cannot create cache dir: %w", err)
		}

		slog.Info(
			"downloading yt-dlp binary",
			"url", "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp",
			"dest", path,
		)
	
		if err := getter.GetAny(cachePath, "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp"); err != nil {
			slog.Error(
				"failed to download yt-dlp",
				"url", "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp",
				"dest", cachePath,
				"err", err,
			)
			return "", fmt.Errorf("failed to download yt-dlp: %w", err)
		}

		if err := os.Chmod(path, 0o700); err != nil {
			slog.Error(
				"failed to make yt-dlp executable",
				"path", path,
				"err", err,
			)
			return "", fmt.Errorf("cannot make yt-dlp executable: %w", err)
		}
	} else {
		slog.Debug(
			"yt-dlp found in cache",
			"path", path,
		)
	}

	return path, nil
}

func ClearYtDlp() error {
	binary := "yt-dlp"
	cachePath, err := getCachePath()
	if err != nil {
		return fmt.Errorf("cannot get cache path: %w", err)
	}
	
	path := filepath.Join(cachePath, binary)

	slog.Info(
		"removing yt-dlp from cache",
		"path", path,
	)

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("unable to remove %s: %w", path, err)
	}
	return nil
}
