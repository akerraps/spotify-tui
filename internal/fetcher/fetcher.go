package fetcher

import (
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"akerraps/tunectl/internal/cache"
	"akerraps/tunectl/internal/types"
)

func songExists(prefix string) (bool, error) {
	matches, err := filepath.Glob(prefix + ".*")
	if err != nil {
		return false, err
	}
	return len(matches) > 0, nil
}

func FetchAudio(tracks []types.TrackInfo, out string) error {
	bin, err := cache.GetYtDlp()
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
			"-t", "mp3",
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
