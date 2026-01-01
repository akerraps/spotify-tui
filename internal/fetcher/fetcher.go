package fetcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"akerraps/tunectl/internal/types"

	"github.com/hashicorp/go-getter"
)

func checkDependencies() (string, error) {
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

		if err := getter.GetAny(cachePath, "https://github.com/yt-dlp/yt-dlp/releases/latest/diownload/yt-dlp"); err != nil {
			return "", fmt.Errorf("failed to download yt-dlp: %w", err)
		}

		if err := os.Chmod(path, 0o700); err != nil {
			return "", fmt.Errorf("cannot make yt-dlp executable: %w", err)
		}
	}

	return path, nil
}

func songExists(prefix string) (bool, error) {
	matches, err := filepath.Glob(prefix + ".*")
	if err != nil {
		return false, err
	}
	return len(matches) > 0, nil
}

func FetchAudio(tracks []types.TrackInfo, out string) error {
	bin, err := checkDependencies()
	if err != nil {
		return err
	}

	for _, song := range tracks {

		name := song.Name
		artist := strings.Join(song.Artists, " ")
		output := filepath.Join(out, name)

		exists, err := songExists(output)
		if err != nil {
			return err
		}

		if exists {
			log.Printf("already exists: %s - %s", name, artist)
			continue
		}

		log.Printf("fetching %s - %s", name, artist)
		cmd := exec.Command(bin,
			"-x",
			"--restrict-filenames",
			"--windows-filenames",
			"--quiet",
			"--no-warnings",
			"-f", "bestaudio",
			"ytsearch:"+name+" "+artist,
			"-o", output)

		_, err = cmd.Output()

		if err != nil {
			log.Printf("failed to fetch %s - %s: %v", name, artist, err)
			continue
		} else {
			log.Printf("fetched %s - %s", name, artist)
		}
	}

	return nil
}
