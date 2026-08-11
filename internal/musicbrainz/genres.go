package musicbrainz

import (
	"akerraps/tunectl/internal/types"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"github.com/michiwend/gomusicbrainz"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func extractTags(tags []gomusicbrainz.Tag, info *types.TrackInfo) {
	c := cases.Title(language.Und)

	for _, tag := range tags {
		tagName := strings.ToLower(strings.TrimSpace(tag.Name))

		if tagName == "" {
			continue
		}

		if isYear(tagName) {
			if info.Year == 0 {
				info.Year, _ = strconv.Atoi(tagName)
			}
			continue
		}

		if slices.Contains(types.GenreFamilies, tagName) {
			genreName := c.String(tagName)

			if genreExists(info.Genres, genreName) {
				continue
			}

			info.Genres = append(info.Genres, genreName)

			slog.Debug(
				"genre added",
				"genre", genreName,
			)

			continue
		}

		if !slices.Contains(info.Tags, tagName) {
			info.Tags = append(info.Tags, tagName)
		}
	}

	slog.Debug(
		"artist tags resolved",
		"genres", info.Genres,
		"year", info.Year,
		"tags", info.Tags,
	)
}

func isYear(value string) bool {
	if len(value) != 4 {
		return false
	}

	for _, r := range value {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

func genreExists(genres []string, candidate string) bool {
	candidate = strings.ToLower(candidate)

	for _, genre := range genres {
		genre = strings.ToLower(genre)

		if genre == candidate {
			return true
		}
	}

	return false
}
