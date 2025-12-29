package fetcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/hashicorp/go-getter"
)

func checkDependencies() (binPath string, err error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	binary := "yt-dlp"
	cachePath := fmt.Sprintf("/home/%s/.cache/tunectl/", currentUser.Username)

	if _, err := os.Stat(cachePath + binary); os.IsNotExist(err) {
		log.Println("Getting yt-dlp dependency")
		if err := getter.GetAny(cachePath, "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp"); err != nil {
			return "", err
		}

		path := filepath.Join(cachePath, "yt-dlp")
		if err := os.Chmod(path, 0o700); err != nil {
			return "", fmt.Errorf("cannot make yt-dlp executable: %w", err)
		}
	}

	return cachePath + binary, nil
}

func FetchAudio(videoID string) error {
	bin, err := checkDependencies()
	if err != nil {
		return err
	}

	cmd := exec.Command(bin, "-h")

	stdout, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(stdout))

	return nil
}
