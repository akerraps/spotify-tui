package musicbrainz

import (
	"log/slog"
	"time"

	"github.com/michiwend/gomusicbrainz"
)

func getArtistInfo(artistID string, genres []string) ([]string, error) {
	slog.Debug(
		"fetching artist information",
		"artist_id", artistID,
	)

	client := createClient()

	var resp *gomusicbrainz.ArtistSearchResponse
	var err error

	for attempt := 1; attempt <= 3; attempt++ {
		resp, err = client.SearchArtist("arid:"+artistID, 1, 0)

		if err == nil {
			break
		}

		slog.Warn(
			"artist lookup failed, retrying",
			"attempt", attempt,
			"max_attempts", 3,
			"artist_id", artistID,
			"err", err,
		)

		time.Sleep(time.Duration(attempt) * time.Second * 5)
	}

	if err != nil {
		slog.Error(
			"artist lookup failed",
			"artist_id", artistID,
			"err", err,
		)

		return nil, err
	}

	tags := resp.Artists[0].Tags

	names := extractGenres(tags, genres)

	return names, nil
}
