package fetcher

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"time"

	"akerraps/tunectl/internal/types"

	"github.com/michiwend/gomusicbrainz"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func createClient() (client *gomusicbrainz.WS2Client) {
	client, _ = gomusicbrainz.NewWS2Client(
		"https://musicbrainz.org/ws/2",
		"Terminal music downloader & metadata tool",
		"0.0.1-beta",
		"https://github.com/akerraps/tunectl")

	return client
}

func GetSongInfo(song string, artist []string) (info types.TrackInfo, err error) {
	slog.Debug(
		"searching musicbrainz recording",
		"song", song,
		"artists", artist,
	)

	client := createClient()

	info = types.TrackInfo{
		Title:   song,
		Artists: artist,
	}

	query := ""
	if len(artist) == 0 {
		query = fmt.Sprintf(`recording:"%s"`, song)
	} else {
		query = fmt.Sprintf(`recording:"%s" AND artist:%s`, song, artist[0])
	}

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
			artist,
			err,
		)
	}

	if len(resp.Recordings) == 0 {
		slog.Warn(
			"no information found for song",
			"song", song,
			"artists", artist,
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

		info.Genres, err = getArtistInfo(artistId)
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

func getArtistInfo(artistID string) ([]string, error) {
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

	names := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagName := strings.ToLower(tag.Name)

		isGenre := slices.Contains(types.GenreFamilies, tagName)

		var exists bool

		if isGenre {
			exists = slices.Contains(names, tagName)
		}

		slog.Debug(
			"processing artist tag",
			"tag", tag.Name,
			"normalized", tagName,
			"is_genre", isGenre,
			"already_added", exists,
		)

		if !isGenre {
			continue
		}

		if !exists {
			names = append(
				names,
				cases.Title(language.English).String(tag.Name),
			)

			slog.Debug(
				"genre added",
				"genre", tag.Name,
			)
		}
	}

	slog.Debug(
		"artist genres resolved",
		"artist_id", artistID,
		"genres", names,
	)

	return names, nil
}
