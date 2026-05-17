package fetcher

import (
	"fmt"
	"log"

	"github.com/michiwend/gomusicbrainz"
)

type Tag struct {
	Count int
	Name  string
}

func createClient() (client *gomusicbrainz.WS2Client) {
	client, _ = gomusicbrainz.NewWS2Client(
		"https://musicbrainz.org/ws/2",
		"Terminal music downloader & metadata tool",
		"0.0.1-beta",
		"https://github.com/akerraps/tunectl")

	return client
}

func GetSongInfo(song string, artist string) (songName string, artistName string, err error) {
	client := createClient()
	query := fmt.Sprintf(`recording:"%s" AND artist:%s`, song, artist)
	artistId := ""

	resp, err := client.SearchRecording(query, 1, 0)
	if err != nil {
		return song, artist, fmt.Errorf("cannot fetch \"%s - %s\" song data: %w", song, artist, err)
	}

	if len(resp.Recordings) == 0 {
		log.Printf("no information found for song \"%s\" and artist \"%s\"\n", song, artist)
	} else {

		song = resp.Recordings[0].Title
		artist = resp.Recordings[0].ArtistCredit.NameCredits[0].Artist.Name
		artistId = string(resp.Recordings[0].ArtistCredit.NameCredits[0].Artist.ID)
		fmt.Println(getArtistInfo(artistId))
	}

	return song, artist, err
}

func getArtistInfo(artistID string) ([]string, error) {
	client := createClient()

	resp, err := client.SearchArtist("arid:"+artistID, 1, 0)
	if err != nil {
		return nil, err
	}

	tags := resp.Artists[0].Tags

	names := make([]string, 0, len(tags))
	for _, tag := range tags {
		names = append(names, tag.Name)
	}

	return names, nil
}
