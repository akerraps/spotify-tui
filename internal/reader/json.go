package reader

import (
	"akerraps/tunectl/internal/types"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func ReadJsonFile(filePath string) (tracks []types.TrackInfo, err error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to access input path %s: %w", filePath, err)
	}

	if info.IsDir() {
		files, err := filepath.Glob(filepath.Join(filePath, "*.json"))
		if err != nil {
			return nil, fmt.Errorf("unable to find JSON files in %s: %w", filePath, err)
		}

		if len(files) == 0 {
			return nil, fmt.Errorf("no JSON files found in %s", filePath)
		}

		for _, file := range files {
			fileTracks, err := readSingleJsonFile(file)
			if err != nil {
				return nil, err
			}

			tracks = append(tracks, fileTracks...)
		}

		slog.Info(
			"json files parsed",
			"files", len(files),
			"tracks", len(tracks),
		)

		return tracks, nil
	}

	// If the path is a file, read it directly.
	return readSingleJsonFile(filePath)
}

func readSingleJsonFile(filePath string) (tracks []types.TrackInfo, err error) {
	slog.Info(
		"reading json file",
		"file", filePath,
	)

	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf(
			"unable to read input file %s: %w",
			filePath,
			err,
		)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&tracks)
	if err != nil {
		return nil, fmt.Errorf(
			"unable to parse JSON %s: %w",
			filePath,
			err,
		)
	}

	slog.Info(
		"json file parsed",
		"file", filePath,
		"tracks", len(tracks),
	)

	return tracks, nil
}
