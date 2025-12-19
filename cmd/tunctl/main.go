package main

import (
	"os"

	"akerraps/tunectl/fetcher"
	"akerraps/tunectl/internal/cli"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	if len(os.Args) > 1 {
		cli.RunCli()
	} else {
		fetcher.FetchAudio("")
		// tui.RunTui()
	}
}
