package main

import (
	"context"
	"fmt"
	"log"

	authenticate "akerraps/spotify-tui/authenticate"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	authenticate.InitAuth()
	client := authenticate.StartLogin()

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
}
