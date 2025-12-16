package core

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

type TrackInfo struct {
	Name    string
	Artists []string
}

func ListPlaylists(ctx context.Context, client *spotify.Client) spotify.SimplePlaylistPage {
	playlists, _ := client.CurrentUsersPlaylists(ctx)
	return *playlists
}

func PlaylistData(ctx context.Context, client *spotify.Client, playlistID spotify.ID) (spotify.FullPlaylist, error) {
	fullPlaylist, err := client.GetPlaylist(ctx, playlistID)
	if err != nil {
		return spotify.FullPlaylist{}, err
	}

	return *fullPlaylist, nil
}

func PlaylistTracks(playlist spotify.FullPlaylist) []TrackInfo {
	results := []TrackInfo{}
	for _, entry := range playlist.Tracks.Tracks {
		artists := []string{}
		for i := range entry.Track.Artists {
			artists = append(artists, entry.Track.Artists[i].Name)
		}
		info := TrackInfo{
			Name:    entry.Track.Name,
			Artists: artists,
		}
		results = append(results, info)
	}
	return results
}
