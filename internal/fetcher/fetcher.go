package fetcher

import (
	"fmt"
	"log/slog"
	"os/exec"
	"path/filepath"
	"strings"

	"akerraps/tunectl/internal/cache"
	"akerraps/tunectl/internal/musicbrainz"
	"akerraps/tunectl/internal/types"

	"go.senan.xyz/taglib"
)

func songExists(path string) (bool, error) {
	prefix := strings.TrimSuffix(path, filepath.Ext(path))

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
		taglib.Title:  {song.Title},
	}, 0)
	if err != nil {
		return err
	} else {
		slog.Debug(
			"metadata writen",
			"title", song.Title,
			"artists", song.Artists,
			"genres", song.Genres,
			"album", song.Album,
		)
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
		slog.Debug(
			"processing track",
			"title", song.Title,
			"artists", song.Artists,
		)

		fileName := strings.ReplaceAll(
			song.Title+"_"+song.Artists[0]+".mp3",
			" ",
			"_",
		)
		output := filepath.Join(opts.OutputDir, fileName)

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

				info, err := musicbrainz.GetSongInfo(song.Title, song.Artists, song.Genres)
				if err != nil {
					slog.Error(
						"failed to fetch metadata, skipping song",
						"song", song.Title,
						"artists", song.Artists,
						"err", err,
					)
					continue
				}

				song.Title = info.Title
				song.Artists = info.Artists
				song.Genres = append(song.Genres, info.Genres...)

				err = writeMetadata(output, song)
				if err != nil {
					slog.Warn(
						"failed to write metadata",
						"title", song.Title,
						"artists", song.Artists,
						"file", output,
						"err", err,
					)
					continue
				}
			}
			continue
		}

		if opts.NoAPI == false {
			info, err := musicbrainz.GetSongInfo(song.Title, song.Artists, song.Genres)
			if err != nil {
				slog.Error(
					"failed to fetch metadata, skipping song",
					"song", song.Title,
					"artists", song.Artists,
					"err", err,
				)
				continue
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
			"--audio-quality", "128K",
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

		err = writeMetadata(output, song)
		if err != nil {
			slog.Warn(
				"failed to write metadata",
				"song", song.Title,
				"artists", song.Artists,
				"err", err,
			)
			continue
		} else {
			slog.Debug(
				"metadata written",
				"title", song.Title,
				"file", output,
			)
		}
	}

	return nil
}
