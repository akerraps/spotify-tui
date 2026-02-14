package core

import (
	"context"
	"fmt"
	"strings"

	"akerraps/tunectl/internal/fetcher"
)

type Service struct {
	Name string
}

func (s *Service) RunPlaylists(ctx context.Context) error {
	client := Auth(ctx)
	myPlaylists := listPlaylists(ctx, client)

	for _, p := range myPlaylists.Playlists {
		fmt.Printf("Playlist found: %s\n", p.Name)
	}
	return nil
}

func (s *Service) RunSongs(ctx context.Context, playlistName string, download bool, out string) error {
	client := Auth(ctx)
	myPlaylists := listPlaylists(ctx, client)

	found := false
	for _, p := range myPlaylists.Playlists {
		if playlistName == p.Name {
			found = true
			playlistID := p.ID
			myPlaylistData, err := playlistData(ctx, client, playlistID)
			if err != nil {
				return err
			}
			myTrackInfo := tracks(myPlaylistData)

			if download {
				fetcher.FetchAudio(myTrackInfo, out)
			} else {
				for _, song := range myTrackInfo {

					name := song.Title
					artist := strings.Join(song.Artists, " ")
					album := song.Album
					albumArtist := strings.Join(song.AlbumArtist, " ")

					fmt.Printf("Song name: %s, Artists: %s, Album: %s, Artis Album: %s\n", name, artist, album, albumArtist)
				}
			}
		}
	}
	if !found {
		return fmt.Errorf("playlist %q not found", playlistName)
	}
	return nil
}
