package playlists

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify/v2"
)

func ListPlaylists(ctx context.Context, client *spotify.Client) {
	playlists, _ := client.CurrentUsersPlaylists(ctx)
	for _, p := range playlists.Playlists {
		fmt.Println("Playlist:", p.Name)
		playlistID := p.ID

		fullPlaylist, _ := client.GetPlaylist(ctx, playlistID)

		fmt.Println(fullPlaylist.Tracks)
	}
}
