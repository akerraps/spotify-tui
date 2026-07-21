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
		tagName := strings.ToLower(tag.Name)

		isGenre := slices.Contains(types.GenreFamilies, tagName)

		var exists bool

		if isGenre {
			exists = slices.Contains(names, c.String(tagName))
			if !exists {
				exists = slices.Contains(existing, c.String(tagName))
			}
		}

		slog.Debug(
			"processing artist tag",
			"tag", tag.Name,
			"normalized", tagName,
			"is_genre", isGenre,
			"already_added", exists,
		)

		if !isGenre {
			continue
		}

		if !exists {
			names = append(
				names,
				c.String(tag.Name),
			)

			slog.Debug(
				"genre added",
				"genre", tag.Name,
			)
		}
	}

	slog.Debug(
		"artist genres resolved",
		"genres", names,
	)

	return names
}
