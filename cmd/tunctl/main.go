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

func (a *App) RunTasks(ctx context.Context) error {
	authenticate.InitAuth()
	client := authenticate.StartLogin()

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

	myPlaylists := playlists.ListPlaylists(ctx, client)

	fmt.Println(myPlaylists)
	return nil
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
					return appCore.RunTasks(c.Context)
				},
			},
		},
	}

	cliApp.RunContext(ctx, os.Args)
}

// func main() {
// 	authenticate.InitAuth()
// 	client := authenticate.StartLogin()
//
// 	ctx := context.Background()
// 	// use the client to make calls that require authorization
// 	user, err := client.CurrentUser(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("You are logged in as:", user.ID)
//
// 	myPlaylists := playlists.ListPlaylists(ctx, client)
//
// 	for _, p := range myPlaylists.Playlists {
// 		playlistID := p.ID
// 		myPlaylistData := playlists.PlaylistData(ctx, client, playlistID)
// 		myTrackInfo := playlists.PlaylistTracks(myPlaylistData)
// 		fmt.Println(myTrackInfo)
// 	}
// }
