package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"akerraps/tunectl/internal/authenticate"
	playlists "akerraps/tunectl/internal/spotify"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func init() {
	godotenv.Load()
}

type App struct {
	Name string
}

func NewApp(ctx context.Context) *App {
	return &App{
		Name: "Tunectl",
	}
}

func main() {
	ctx := context.Background()

	appCore := NewApp(ctx)

	cliApp := &cli.App{
		Name: "Tunectl",
		Commands: []*cli.Command{
			{
				Name: "playlists",
				Action: func(c *cli.Context) error {
					return appCore.RunPlaylists(c.Context)
				},
			},
			{
				Name:      "songs",
				Usage:     "Lists songs from a specefied playlist",
				ArgsUsage: "<playlist>",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return fmt.Errorf("you must specify a playlist")
					}

					playlist := c.Args().Get(0)

					return appCore.RunSogns(c.Context, playlist)
				},
			},
		},
	}
	err := cliApp.RunContext(ctx, os.Args)
	if err != nil {
		log.Fatal(err)
	}
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
