package core

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"

	"akerraps/tunectl/internal/types"
)

func ExportTracksToCSV(tracks []types.TrackInfo, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	err = writer.Write([]string{"Title", "Artists", "Album", "AlbumArtist"})
	if err != nil {
		return err
	}

	for _, t := range tracks {
		err := writer.Write([]string{
			t.Title,
			strings.Join(t.Artists, ", "),
			t.Album,
			strings.Join(t.AlbumArtist, ", "),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func ExportTracksToJSON(tracks []types.TrackInfo, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(tracks)
}
