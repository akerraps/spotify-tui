package core

import (
	"context"
	"fmt"
)

type Service struct {
	Name string
}

func (a *Service) RunPlaylists(ctx context.Context) error {
	client := Auth(ctx)
	myPlaylists := ListPlaylists(ctx, client)

	for _, p := range myPlaylists.Playlists {
		fmt.Printf("Playlist found: %s\n", p.Name)
	}
	return nil
}

func (a *Service) RunSogns(ctx context.Context, playlistName string) error {
	client := Auth(ctx)
	myPlaylists := ListPlaylists(ctx, client)

	found := false
	for _, p := range myPlaylists.Playlists {
		if playlistName == p.Name {
			found = true
			playlistID := p.ID
			myPlaylistData, err := PlaylistData(ctx, client, playlistID)
			if err != nil {
				return err
			}
			myTrackInfo := PlaylistTracks(myPlaylistData)
			fmt.Println(myTrackInfo)
		}
	}
	if !found {
		return fmt.Errorf("playlist %q not found", playlistName)
	}
	return nil
}
