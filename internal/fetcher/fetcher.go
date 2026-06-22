package fetcher

import (
	"fmt"
	"log"
	"log/slog"
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
	} else {
		slog.Debug(
			"using yt-dlp binary",
			"path", bin,
		)
	}

	for _, song := range tracks {

		output := filepath.Join(opts.OutputDir, song.Title)

		exists, err := songExists(output)
		if err != nil {
			return err
		}

		if exists {
			slog.Info(
				"song already exists",
				"title", song.Title,
				"artists", song.Artists,
			)

			if opts.NoAPI == false {
				slog.Debug(
					"fetching metadata from api",
					"title", song.Title,
					"artists", song.Artists,
				)

				info, err := GetSongInfo(song.Title, song.Artists)
				if err != nil {
					return err
				}

				song.Title = info.Title
				song.Artists = info.Artists
				song.Genres = append(song.Genres, info.Genres...)

				err = writeMetadata(output+".mp3", song)
				if err != nil {
					slog.Warn(
						"failed to write metadata",
						"title", song.Title,
						"artists", song.Artists,
						"file", output+".mp3",
						"err", err,
					)
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

		slog.Debug(
			"starting download",
			"title", song.Title,
			"artists", song.Artists,
			"query", search,
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
			slog.Error(
				"download failed",
				"title", song.Title,
				"artists", song.Artists,
				"err", err,
			)
			continue
		} else {
			slog.Info(
				"download completed",
				"title", song.Title,
				"artists", song.Artists,
			)
		}

		err = writeMetadata(output+".mp3", song)
		if err != nil {
			log.Printf("coudnt write metadata to %s - %s: %v", song.Title, song.Artists, err)
			continue
		} else {
			slog.Debug(
				"metadata written",
				"title", song.Title,
				"file", output+".mp3",
			)
		}
	}

	return nil
}
