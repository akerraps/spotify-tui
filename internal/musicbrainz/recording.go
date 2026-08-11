package musicbrainz

import (
	"akerraps/tunectl/internal/types"
	"fmt"
	"log/slog"
	"time"

	"github.com/michiwend/gomusicbrainz"
)

func searchQuery(song string, artists []string) (query string) {
	if len(artists) == 0 {
		query = fmt.Sprintf(`recording:"%s"`, song)
	} else {
		query = fmt.Sprintf(`recording:"%s" AND artist:%s`, song, artists[0])
	}

	return query
}

func enrichTrack(info *types.TrackInfo, resp *gomusicbrainz.RecordingSearchResponse) error {
	recording := resp.Recordings[0]

	info.Artists = info.Artists[:0]
	info.Title = recording.Title

	for _, credit := range recording.ArtistCredit.NameCredits {
		info.Artists = append(info.Artists, credit.Artist.Name)
	}

	info.ArtistID = string(recording.ArtistCredit.NameCredits[0].Artist.ID)

	if err := getArtistInfo(info); err != nil {
		return err
	}

	slog.Debug(
		"metadata enriched",
		"title", info.Title,
		"artists", info.Artists,
		"album", info.Album,
		"genres", info.Genres,
		"year", info.Year,
		"tags", info.Tags,
	)

	return nil
}

func GetSongInfo(song *types.TrackInfo) error {
	slog.Debug(
		"searching musicbrainz recording",
		"song", song.Title,
		"artists", song.Artists,
	)

	client := createClient()

	query := searchQuery(song.Title, song.Artists)

	var resp *gomusicbrainz.RecordingSearchResponse
	var err error

	for attempt := 1; attempt <= 3; attempt++ {
		resp, err = client.SearchRecording(query, 1, 0)

		if err == nil {
			break
		}

		slog.Warn(
			"musicbrainz recording search failed, retrying",
			"attempt", attempt,
			"err", err,
			"song", song.Title,
		)

		time.Sleep(time.Duration(attempt) * 5 * time.Second)
	}

	if err != nil {
		return fmt.Errorf(
			`cannot fetch "%s - %v" song data: %w`,
			song.Title,
			song.Artists,
			err,
		)
	}

	if len(resp.Recordings) == 0 {
		slog.Warn(
			"no information found for song",
			"song", song.Title,
			"artists", song.Artists,
		)

		return nil
	}

	slog.Debug(
		"recording found",
		"title", resp.Recordings[0].Title,
		"artists", song.Artists,
	)

	if err := enrichTrack(song, resp); err != nil {
		return err
	}

	return nil
}
