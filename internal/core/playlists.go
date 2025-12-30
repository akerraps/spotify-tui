package core

import (
	"context"

	"akerraps/tunectl/internal/types"

	"github.com/zmb3/spotify/v2"
)

// Get a list of playlists; used for listing
func listPlaylists(ctx context.Context, client *spotify.Client) spotify.SimplePlaylistPage {
	playlists, _ := client.CurrentUsersPlaylists(ctx)
	return *playlists
}

// Get data from a certain plalist; returns FullPlaylist type data from spotify wrapper which will be used in tracks()
func playlistData(ctx context.Context, client *spotify.Client, playlistID spotify.ID) (spotify.FullPlaylist, error) {
	fullPlaylist, err := client.GetPlaylist(ctx, playlistID)
	if err != nil {
		return spotify.FullPlaylist{}, err
	}

	return *fullPlaylist, nil
}

func tracks(playlist spotify.FullPlaylist) []types.TrackInfo {
	results := []types.TrackInfo{}
	for _, entry := range playlist.Tracks.Tracks {
		artists := []string{}
		for i := range entry.Track.Artists {
			artists = append(artists, entry.Track.Artists[i].Name)
		}
		info := types.TrackInfo{
			Name:    entry.Track.Name,
			Artists: artists,
		}
		results = append(results, info)
	}
	return results
}
