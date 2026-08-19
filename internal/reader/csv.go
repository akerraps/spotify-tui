package reader

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"akerraps/tunectl/internal/types"
)

func ReadCsvFile(path string) ([]types.TrackInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("unable to access input path %s: %w", path, err)
	}

	if info.IsDir() {
		files, err := filepath.Glob(filepath.Join(path, "*.csv"))
		if err != nil {
			return nil, fmt.Errorf("unable to find CSV files in %s: %w", path, err)
		}

		if len(files) == 0 {
			return nil, fmt.Errorf("no CSV files found in %s", path)
		}

		var tracks []types.TrackInfo

		for _, file := range files {
			fileTracks, err := readSingleCsvFile(file)
			if err != nil {
				return nil, err
			}

			tracks = append(tracks, fileTracks...)
		}

		slog.Info(
			"csv files parsed",
			"files", len(files),
			"tracks", len(tracks),
		)

		return tracks, nil
	}

	return readSingleCsvFile(path)
}

func readSingleCsvFile(filePath string) (tracks []types.TrackInfo, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s: %w", filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s: %w", filePath, err)
	}

	if len(records) == 0 {
		slog.Warn(
			"csv file is empty",
			"file", filePath,
		)
		return nil, nil
	}

	// Skip header.
	records = records[1:]

	slog.Debug(
		"csv loaded",
		"file", filePath,
		"rows", len(records),
	)

	for _, row := range records {
		if len(row) < 4 {
			slog.Warn(
				"skipping invalid csv row",
				"file", filePath,
				"columns", len(row),
			)
			continue
		}

		song := types.TrackInfo{
			Title: row[0],
			Album: row[2],
		}

		artists := strings.Split(row[1], ",")
		for i := range artists {
			artists[i] = strings.TrimSpace(artists[i])
		}
		song.Artists = artists

		genres := strings.Split(row[3], ",")
		for i := range genres {
			genres[i] = strings.TrimSpace(genres[i])
		}
		song.Genres = genres

		tracks = append(tracks, song)
	}

	slog.Info(
		"csv file parsed",
		"file", filePath,
		"tracks", len(tracks),
	)

	return tracks, nil
}
