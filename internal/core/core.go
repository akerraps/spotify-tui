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

	myPlaylists, err := listPlaylists(ctx, client)
	if err != nil {
		return err
	}

	for _, p := range myPlaylists {
		fmt.Printf("Playlist found: %s\n", p.Name)
	}

	return nil
}

func (s *Service) RunSongs(ctx context.Context, playlistName string, download bool, out string) error {
	client := Auth(ctx)

	myPlaylists, err := listPlaylists(ctx, client)
	if err != nil {
		return err
	}

	for _, p := range myPlaylists {
		if playlistName == p.Name {

			playlistID := p.ID

			myPlaylistData, err := playlistData(ctx, client, playlistID)
			if err != nil {
				return err
			}

			myTrackInfo, err := tracks(ctx, client, myPlaylistData)
			if err != nil {
				return err
			}

			if download {
				fetcher.FetchAudio(myTrackInfo, out)
				return nil
			}

			if out != "" {
				if strings.HasSuffix(out, ".json") {
					err = ExportTracksToJSON(myTrackInfo, out)
				} else {
					err = ExportTracksToCSV(myTrackInfo, out)
				}

				if err != nil {
					return err
				}

				fmt.Println("Exported to", out)
				return nil
			}

			fmt.Printf("         Name        |         Artist       |              Album             |    Album Artist     \n")

			for _, song := range myTrackInfo {
				name := song.Title
				artist := strings.Join(song.Artists, ", ")
				album := song.Album
				albumArtist := strings.Join(song.AlbumArtist, ", ")

				fmt.Printf(
					"%-20s | %-20s | %-30s | %-20s\n",
					name, artist, album, albumArtist,
				)
			}

			return nil
		}
	}

	return fmt.Errorf("playlist %q not found", playlistName)
}
