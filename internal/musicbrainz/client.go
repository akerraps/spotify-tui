package musicbrainz

import "github.com/michiwend/gomusicbrainz"

func createClient() (client *gomusicbrainz.WS2Client) {
	client, _ = gomusicbrainz.NewWS2Client(
		"https://musicbrainz.org/ws/2",
		"Terminal music downloader & metadata tool",
		"0.0.1-beta",
		"https://github.com/akerraps/tunectl")

	return client
}
