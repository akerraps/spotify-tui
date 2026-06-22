package fetcher

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"akerraps/tunectl/internal/cache"
	"akerraps/tunectl/internal/types"

	"go.senan.xyz/taglib"
)

func songExists(prefix string) (bool, error) {
	matches, err := filepath.Glob(prefix + ".*")
	if err != nil {
		return false, err
	}
	return len(matches) > 0, nil
}

func writeMetadata(file string, song types.TrackInfo) error {

	err := taglib.WriteTags(file, map[string][]string{
		taglib.Album:  {song.Album},
		taglib.Artist: song.Artists,
		taglib.Genre:  song.Genres,
	}, 0)
	if err != nil {
		return err
	}
	return nil
}

func FetchAudio(tracks []types.TrackInfo, opts types.Options) error {
	bin, err := cache.GetYtDlp()
	if err != nil {
		return err
	}

	for _, song := range tracks {

		log.Printf("fetching %s - %s", song.Title, strings.Join(song.Artists, ", "))

		output := filepath.Join(opts.OutputDir, song.Title)

		exists, err := songExists(output)
		if err != nil {
			return err
		}

		if exists {
			log.Printf("already exists: %s - %s", song.Title, strings.Join(song.Artists, ", "))

			if opts.NoAPI == false {
				info, err := GetSongInfo(song.Title, song.Artists)
				if err != nil {
					return err
				}

				song.Title = info.Title
				song.Artists = info.Artists
				song.Genres = append(song.Genres, info.Genres...)

				err = writeMetadata(output+".mp3", song)
				if err != nil {
					log.Printf("coudnt write metadata to %s - %s: %v", song.Title, strings.Join(song.Artists, ", "), err)
					continue
				}
			}
			continue
		}

		if opts.NoAPI == false {
			info, err := GetSongInfo(song.Title, song.Artists)
			if err != nil {
				return err
			}

			song.Title = info.Title
			song.Artists = info.Artists
			song.Genres = append(song.Genres, info.Genres...)
		}

		artists := strings.Join(song.Artists, " ")
		search := fmt.Sprintf(
			`ytsearch:"%s" %s "song"`,
			song.Title,
			artists,
		)

		cmd := exec.Command(bin,
			"-x",
			"--restrict-filenames",
			"--quiet",
			"--no-warnings",
			"--embed-thumbnail",
			"--no-playlist",
			"-t", "mp3",
			"--audio-quality", "0",
			search,
			"-o", output,
		)

		_, err = cmd.Output()

		if err != nil {
			log.Printf("failed to fetch %s - %s: %v", song.Title, strings.Join(song.Artists, ", "), err)
			continue
		} else {
			log.Printf("fetched %s - %s", song.Title, strings.Join(song.Artists, ", "))
		}

		err = writeMetadata(output+".mp3", song)
		if err != nil {
			log.Printf("coudnt write metadata to %s - %s: %v", song.Title, strings.Join(song.Artists, ", "), err)
			continue
		}
	}

	return nil
}
