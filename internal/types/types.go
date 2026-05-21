package types

type TrackInfo struct {
	Title   string
	Artists []string

	Album       string
	AlbumArtist []string
}

var GenreFamilies = []string{
	"rock",
	"punk",

	"metal",

	"electronic",
	"edm",

	"hip hop",
	"rap",
	"trap",

	"pop",

	"jazz",

	"classical",
	"opera",
	"contemporary classical",

	"blues",
	"rhythm and blues",
	"r&b",
	"soul",
	"funk",
	"gospel",

	"latin",
	"reggaeton",
	"salsa",
	"bachata",
	"cumbia",
	"flamenco",

	"reggae",
	"dub",
	"dancehall",
	"ska",

	"folk",
	"country",
	"americana",
	"bluegrass",
	"celtic",

	"ambient",
	"experimental",
	"noise",
	"industrial",
	"drone",

	"house",
	"techno",
	"trance",
	"hardcore",
	"breakbeat",
}
