package fetcher

import (
	"fmt"

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

func GetArtistById() {

	client := createClient()
	artist, err := client.LookupArtist("")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", artist)
}
