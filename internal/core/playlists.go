package core

import (
	"context"

	"akerraps/tunectl/internal/types"

	"github.com/zmb3/spotify/v2"
)

// Get a list of playlists; used for listing
func listPlaylists(ctx context.Context, client *spotify.Client) ([]spotify.SimplePlaylist, error) {
	page, err := client.CurrentUsersPlaylists(ctx)
	if err != nil {
		return nil, err
	}

	var playlists []spotify.SimplePlaylist

	for {
		playlists = append(playlists, page.Playlists...)

		err = client.NextPage(ctx, page)
		if err != nil {
			break
		}
	}

	return playlists, nil
}

// Get data from a certain plalist; returns FullPlaylist type data from spotify wrapper which will be used in tracks()
func playlistData(ctx context.Context, client *spotify.Client, playlistID spotify.ID) (spotify.FullPlaylist, error) {
	fullPlaylist, err := client.GetPlaylist(ctx, playlistID)
	if err != nil {
		return spotify.FullPlaylist{}, err
	}

	return *fullPlaylist, nil
}

func tracks(ctx context.Context, client *spotify.Client, playlist spotify.FullPlaylist) ([]types.TrackInfo, error) {
	results := []types.TrackInfo{}
	page := &playlist.Tracks

	for {
		for _, entry := range page.Tracks {

			track := entry.Track

			artists := []string{}
			albumArtists := []string{}

			for _, a := range track.Artists {
				artists = append(artists, a.Name)
			}

			for _, a := range track.Album.Artists {
				albumArtists = append(albumArtists, a.Name)
			}

			info := types.TrackInfo{
				Title:       track.Name,
				Artists:     artists,
				Album:       track.Album.Name,
				AlbumArtist: albumArtists,
			}

			results = append(results, info)
		}

		err := client.NextPage(ctx, page)
		if err != nil {
			break
		}
	}

	return results, nil
}
