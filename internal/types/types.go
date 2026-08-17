package types

import (
	"fmt"
	"os/user"
)

type TrackInfo struct {
	ID int

	ArtistID string

	Title   string
	Artists []string

	Album string

	Genres []string
}

type Options struct {
	OutputDir   string
	NoAPI       bool
	Parallelism int
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
