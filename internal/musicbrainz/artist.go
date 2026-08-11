package musicbrainz

import (
	"akerraps/tunectl/internal/types"
	"log/slog"
	"time"

	"github.com/michiwend/gomusicbrainz"
)

func getArtistInfo(info *types.TrackInfo) error {
	slog.Debug(
		"fetching artist information",
		"artist_id", info.ArtistID,
	)

	client := createClient()

	var resp *gomusicbrainz.ArtistSearchResponse
	var err error

	for attempt := 1; attempt <= 3; attempt++ {
		resp, err = client.SearchArtist("arid:"+info.ArtistID, 1, 0)

		if err == nil {
			break
		}

		slog.Warn(
			"artist lookup failed, retrying",
			"attempt", attempt,
			"max_attempts", 3,
			"artist_id", info.ArtistID,
			"err", err,
		)

		time.Sleep(time.Duration(attempt) * 5 * time.Second)
	}

	if err != nil {
		slog.Error(
			"artist lookup failed",
			"artist_id", info.ArtistID,
			"err", err,
		)

		return err
	}

	extractTags(resp.Artists[0].Tags, info)

	return nil
}
