package musicbrainz

import (
	"akerraps/tunectl/internal/types"
	"log/slog"
	"slices"
	"strings"

	"github.com/michiwend/gomusicbrainz"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func extractGenres(tags []gomusicbrainz.Tag, existing []string) []string {

	c := cases.Title(language.Und) //Undefined lang.
	names := make([]string, 0, len(tags))

	for _, tag := range tags {
		tagName := strings.ToLower(strings.TrimSpace(tag.Name))

		isGenre := slices.Contains(types.GenreFamilies, tagName)
		if !isGenre {
			continue
		}

		genreName := c.String(tagName)

		exists := genreExists(names, genreName) ||
			genreExists(existing, genreName)

		slog.Debug(
			"processing artist tag",
			"tag", tag.Name,
			"normalized", genreName,
			"is_genre", isGenre,
			"already_added", exists,
		)

		if exists {
			continue
		}

		names = append(names, genreName)

		slog.Debug(
			"genre added",
			"genre", genreName,
		)
	}

	slog.Debug(
		"artist genres resolved",
		"genres", names,
	)

	return names
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
