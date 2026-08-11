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

func enrichTrack(info *types.TrackInfo, resp *gomusicbrainz.RecordingSearchResponse, genres []string) error {
	recording := resp.Recordings[0]

	info.Artists = info.Artists[:0]
	info.Title = recording.Title

	for _, credit := range recording.ArtistCredit.NameCredits {
		info.Artists = append(info.Artists, credit.Artist.Name)
	}

	artistId := string(recording.ArtistCredit.NameCredits[0].Artist.ID)

	genres, err := getArtistInfo(artistId, genres)
	if err != nil {
		return err
	}

	info.Genres = genres

	slog.Debug(
		"metadata enriched",
		"title", info.Title,
		"artists", info.Artists,
		"album", info.Album,
		"genres", info.Genres,
	)

	return nil
}

func GetSongInfo(song types.TrackInfo) (info types.TrackInfo, err error) {
	slog.Debug(
		"searching musicbrainz recording",
		"song", song.Title,
		"artists", song.Artists,
	)

	client := createClient()

	info = song

	query := searchQuery(song.Title, song.Artists)

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
			"song", song.Title,
		)

		time.Sleep(time.Duration(attempt) * time.Second * 5)
	}

	if err != nil {
		return info, fmt.Errorf(
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
		return info, nil
	}

	slog.Debug(
		"recording found",
		"title", info.Title,
		"artists", info.Artists,
	)

	if err := enrichTrack(&info, resp, song.Genres); err != nil {
		return info, err
	}

	return info, nil
}
