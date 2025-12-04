package playlists

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

func ListPlaylists(ctx context.Context, client *spotify.Client) spotify.SimplePlaylistPage {
	playlists, _ := client.CurrentUsersPlaylists(ctx)
	return *playlists
}

func PlaylistData(ctx context.Context, client *spotify.Client, playlistID spotify.ID) spotify.FullPlaylist {
	fullPlaylist, _ := client.GetPlaylist(ctx, playlistID)
	return *fullPlaylist
}
