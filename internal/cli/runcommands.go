package runcommands

import (
	"context"
	"fmt"
)

func (a *App) RunPlaylists(ctx context.Context) error {
	client := auth(ctx)
	myPlaylists := playlists.ListPlaylists(ctx, client)

	for _, p := range myPlaylists.Playlists {
		fmt.Printf("Playlist found: %s\n", p.Name)
	}
	return nil
}

func (a *App) RunSogns(ctx context.Context, playlistName string) error {
	client := auth(ctx)
	myPlaylists := playlists.ListPlaylists(ctx, client)

	for _, p := range myPlaylists.Playlists {
		if playlistName == p.Name {
			playlistID := p.ID
			myPlaylistData := playlists.PlaylistData(ctx, client, playlistID)
			myTrackInfo := playlists.PlaylistTracks(myPlaylistData)
			fmt.Println(myTrackInfo)
		}
	}
	return nil
}
