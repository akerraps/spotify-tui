package fileReader

import (
	"akerraps/tunectl/internal/types"
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func ReadCsvFile(filePath string) (tracks []types.TrackInfo, err error) {
	slog.Info(
		"reading csv file",
		"file", filePath,
	)

	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s: %w", filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	slog.Debug(
		"csv loaded",
		"file", filePath,
		"rows", len(records)-1,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s: %w", filePath, err)
	}

	records = records[1:]

	for _, row := range records {
		if len(row) < 4 {
			slog.Warn(
				"skipping invalid csv row",
				"columns", len(row),
			)
			continue
		}

		song := types.TrackInfo{}

		song.Title = row[0]

		artists := strings.Split(row[1], ",")
		for i := range artists {
			artists[i] = strings.TrimSpace(artists[i])
		}
		song.Artists = artists

		song.Album = row[2]

		genres := strings.Split(row[3], ",")
		for i := range genres {
			genres[i] = strings.TrimSpace(genres[i])
		}
		song.Genres = genres

		tracks = append(tracks, song)
	}

	slog.Info(
		"csv parsed",
		"file", filePath,
		"tracks", len(tracks),
	)

	return tracks, nil
}
