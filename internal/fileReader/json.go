package fileReader

import (
	"akerraps/tunectl/internal/types"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

func ReadJsonFile(filePath string) (tracks []types.TrackInfo, err error) {
	slog.Info(
		"reading json file",
		"file", filePath,
	)

	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s: %w", filePath, err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&tracks)
	if err != nil {
		return nil, fmt.Errorf("unable to parse JSON %s: %w", filePath, err)
	}

	slog.Info(
		"json parsed",
		"file", filePath,
		"tracks", len(tracks),
	)

	return tracks, nil
}
