package cache

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-getter"
)

func getYtDlp(cachePath string) (string, error) {
	binary := "yt-dlp"
	path := filepath.Join(cachePath, binary)

	_, err := os.Stat(path)

	if err == nil {
		slog.Debug(
			"yt-dlp found in cache",
			"path", path,
		)

		return path, nil
	}

	if !os.IsNotExist(err) {
		return "", fmt.Errorf("cannot check yt-dlp: %w", err)
	}

	slog.Info(
		"downloading yt-dlp binary",
		"url", "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp",
		"dest", path,
	)

	if err := getter.GetAny(
		cachePath,
		"https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp",
	); err != nil {
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
