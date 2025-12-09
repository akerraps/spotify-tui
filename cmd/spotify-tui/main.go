package main

import (
	"context"
	"fmt"
	"log"

	authenticate "akerraps/spotify-tui/internal/authenticate"
	playlists "akerraps/spotify-tui/internal/spotify"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	authenticate.InitAuth()
	client := authenticate.StartLogin()

	ctx := context.Background()
	// use the client to make calls that require authorization
	user, err := client.CurrentUser(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

	myPlaylists := playlists.ListPlaylists(ctx, client)

	for _, p := range myPlaylists.Playlists {
		playlistID := p.ID
		myPlaylistData := playlists.PlaylistData(ctx, client, playlistID)
		playlists.PlaylistTracks(myPlaylistData)
	}
}
