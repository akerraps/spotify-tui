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

func GetSongInfo(song string, artists []string, genres []string) (info types.TrackInfo, err error) {
	slog.Debug(
		"searching musicbrainz recording",
		"song", song,
		"artists", artists,
	)

	client := createClient()

	info = types.TrackInfo{
		Title:   song,
		Artists: artists,
	}

	query := searchQuery(song, artists)

	var resp *gomusicbrainz.RecordingSearchResponse

	for attempt := 1; attempt <= 3; attempt++ {
		resp, err = client.SearchRecording(query, 1, 0)

		if err == nil {
			break
		}

		slog.Warn(
			"musicbrainz recording search failed, retrying",
			"attempt", attempt,
			"err", err,
			"song", song,
		)

		time.Sleep(time.Duration(attempt) * time.Second * 5)
	}

	if err != nil {
		return info, fmt.Errorf(
			"cannot fetch \"%s - %s\" song data: %w",
			song,
			artists,
			err,
		)
	}

	if len(resp.Recordings) == 0 {
		slog.Warn(
			"no information found for song",
			"song", song,
			"artists", artists,
		)
	} else {
		slog.Debug(
			"recording found",
			"title", info.Title,
			"artists", info.Artists,
		)

		info.Artists = info.Artists[:0]
		info.Title = resp.Recordings[0].Title
		for _, credit := range resp.Recordings[0].ArtistCredit.NameCredits {
			info.Artists = append(info.Artists, credit.Artist.Name)
		}
		artistId := string(resp.Recordings[0].ArtistCredit.NameCredits[0].Artist.ID)

		info.Genres, err = getArtistInfo(artistId, genres)
		if err != nil {
			return info, err
		}

		slog.Debug(
			"metadata enriched",
			"title", info.Title,
			"artists", info.Artists,
			"genres", info.Genres,
		)

	}

	return info, err
}
