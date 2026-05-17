package fetcher

import (
	"fmt"
	"log"

	"github.com/michiwend/gomusicbrainz"
)

func createClient() (client *gomusicbrainz.WS2Client) {
	client, _ = gomusicbrainz.NewWS2Client(
		"https://musicbrainz.org/ws/2",
		"Terminal music downloader & metadata tool",
		"0.0.1-beta",
		"https://github.com/akerraps/tunectl")

	return client
}

func GetSongInfo(song string, artist string) (string, string, error) {
	client := createClient()
	query := fmt.Sprintf(`recording:"%s" AND artist:%s`, song, artist)

	resp, err := client.SearchRecording(query, 1, 0)
	if err != nil {
		return song, artist, fmt.Errorf("cannot fetch \"%s - %s\" song data: %w", song, artist, err)
	}

	if len(resp.Recordings) == 0 {
		log.Printf("no information found por song \"%s\" and artist \"%s\"\n", song, artist)
	} else {

		song = resp.Recordings[0].Title
		artist = resp.Recordings[0].ArtistCredit.NameCredits[0].Artist.Name

	}

	return song, artist, err
}
