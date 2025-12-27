package main

import (
	"log"
	"os"

	"akerraps/tunectl/internal/cli"
	"akerraps/tunectl/internal/fetcher"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	if len(os.Args) > 1 {
		cli.RunCli()
	} else {
		err := fetcher.FetchAudio("")
		if err != nil {
			log.Fatal(err)
		}
		// fetcher.FetchAudio("")
		// tui.RunTui()
	}
}
