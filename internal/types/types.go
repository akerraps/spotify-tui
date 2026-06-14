package types

import (
	"fmt"
	"os/user"
)

type TrackInfo struct {
	Title   string
	Artists []string

	Album       string
	AlbumArtist []string

	Genres []string
}

type Options struct {
	OutputDir string
	NoAPI     bool
}

func DefaultOptions() Options {
	currentUser, _ := user.Current()

	ouputDir := fmt.Sprintf("/home/%s/Music", currentUser.Username)

	return Options{
		OutputDir: ouputDir,
		NoAPI:     false,
	}
}

var GenreFamilies = []string{
	// Rock family
	"rock",
	"alternative",
	"indie",
	"punk",
	"grunge",
	"shoegaze",
	"emo",
	"post punk",
	"new wave",
	"psych",
	"oi!",

	// Metal family
	"metal",
	"metalcore",
	"deathcore",

	// Pop family
	"pop",
	"synth",
	"dream",

	// Electronic / EDM
	"electronic",
	"edm",
	"house",
	"techno",
	"trance",
	"dubstep",
	"dnb",
	"drum and bass",
	"jungle",
	"breakbeat",
	"ambient",
	"lofi",
	"synthwave",
	"vaporwave",
	"future bass",
	"electro",
	"idm",
	"glitch",

	// Urban
	"hip hop",
	"rap",
	"trap",
	"drill",
	"r&b",
	"soul",
	"funk",

	// Jazz / Blues / Classical
	"jazz",
	"blues",
	"classical",
	"opera",

	// Folk / Country
	"folk",
	"country",
	"americana",
	"bluegrass",
	"celtic",

	// Latin / global
	"latin",
	"reggaeton",
	"salsa",
	"bachata",
	"cumbia",
	"flamenco",

	// Reggae / Caribbean
	"reggae",
	"dub",
	"dancehall",
	"ska",

	// African / modern global
	"afro",
	"afrobeats",
	"amapiano",

	// Experimental
	"experimental",
	"noise",
	"drone",
	"industrial",
}
