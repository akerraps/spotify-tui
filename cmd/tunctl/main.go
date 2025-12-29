package main

import (
	"os"

	"akerraps/tunectl/internal/cli"
	"akerraps/tunectl/internal/tui"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	if len(os.Args) > 1 {
		cli.RunCli()
	} else {
		tui.RunTui()
	}
}
