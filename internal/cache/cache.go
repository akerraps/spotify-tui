package cache

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/hashicorp/go-getter"
)

func GetYtDlp() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("cannot get current user: %w", err)
	}

	binary := "yt-dlp"
	cachePath := fmt.Sprintf("/home/%s/.cache/tunectl/", currentUser.Username)
	path := filepath.Join(cachePath, binary)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("Getting yt-dlp dependency")

		if err := os.MkdirAll(cachePath, 0o755); err != nil {
			return "", fmt.Errorf("cannot create cache dir: %w", err)
		}

		if err := getter.GetAny(cachePath, "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp"); err != nil {
			return "", fmt.Errorf("failed to download yt-dlp: %w", err)
		}

		if err := os.Chmod(path, 0o700); err != nil {
			return "", fmt.Errorf("cannot make yt-dlp executable: %w", err)
		}
	}

	return path, nil
}

func ClearYtDlp() error {
	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("cannot get current user: %w", err)
	}

	binary := "yt-dlp"
	cachePath := fmt.Sprintf("/home/%s/.cache/tunectl/", currentUser.Username)
	path := filepath.Join(cachePath, binary)

	log.Printf("removing %s", path)

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("unable to remove %s: %w", path, err)
	}
	return nil
}
