package fetcher

import (
	"akerraps/tunectl/internal/types"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadCsvFile(filePath string) (tracks []types.TrackInfo, err error) {
	log.Println("Reading CSV:", filePath)
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

	records = records[1:]

	for _, row := range records {
		if len(row) < 3 {
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
		tracks = append(tracks, song)
	}

	return tracks, nil
}

func ReadJsonFile(filePath string) (tracks []types.TrackInfo, err error) {
	log.Println("Reading JSON:", filePath)

	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s: %w", filePath, err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&tracks)
	if err != nil {
		return nil, fmt.Errorf("unable to parse JSON %s: %w", filePath, err)
	}

	return tracks, nil
}
