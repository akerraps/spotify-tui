package types

import (
	"fmt"
	"os/user"
)

type TrackInfo struct {
	ID int

	MBID     string
	ArtistID string

	Title   string
	Artists []string
	Album   string

	Genres []string
}

type Options struct {
	OutputDir   string `envconfig:"OUTPUT_DIR"`
	NoAPI       bool   `envconfig:"NO_API"`
	Parallelism int    `envconfig:"PARALLELISM"`
}

func DefaultOptions() Options {
	currentUser, _ := user.Current()

	ouputDir := fmt.Sprintf("/home/%s/Music", currentUser.Username)

	return Options{
		OutputDir:   ouputDir,
		NoAPI:       false,
		Parallelism: 1,
	}
}
