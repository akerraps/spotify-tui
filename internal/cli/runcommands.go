package runcommands

import (
	"context"
	"fmt"

	"akerraps/tunectl/internal/authenticate"
	playlists "akerraps/tunectl/internal/spotify"
)

type App struct {
	Name string
}

func (a *App) RunPlaylists(ctx context.Context) error {
	client := authenticate.Auth(ctx)
	myPlaylists := playlists.ListPlaylists(ctx, client)

	for _, p := range myPlaylists.Playlists {
		fmt.Printf("Playlist found: %s\n", p.Name)
	}
	return nil
}

func (a *App) RunSogns(ctx context.Context, playlistName string) error {
	client := authenticate.Auth(ctx)
	myPlaylists := playlists.ListPlaylists(ctx, client)

	found := false
	for _, p := range myPlaylists.Playlists {
		if playlistName == p.Name {
			found = true
			playlistID := p.ID
			myPlaylistData, err := playlists.PlaylistData(ctx, client, playlistID)
			if err != nil {
				return err
			}
			myTrackInfo := playlists.PlaylistTracks(myPlaylistData)
			fmt.Println(myTrackInfo)
		}
	}
	if !found {
		return fmt.Errorf("playlist %q not found", playlistName)
	}
	return nil
}
