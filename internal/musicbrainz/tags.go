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

func extractTags(tags []gomusicbrainz.Tag, info *types.TrackInfo) {
	log := slog.With("track_id", info.ID)

	c := cases.Title(language.Und)

	for _, tag := range tags {
		tagName := strings.ToLower(strings.TrimSpace(tag.Name))

		if tagName == "" {
			continue
		}

		genreName := c.String(tagName)

		if valueExists(info.Genres, genreName) {
			continue
		}

		info.Genres = append(info.Genres, genreName)

		log.Debug(
			"genre added",
			"genre", genreName,
		)
	}

	log.Debug(
		"artist tags resolved",
		"genres", info.Genres,
	)
}

func valueExists(values []string, candidate string) bool {
	candidate = strings.ToLower(strings.TrimSpace(candidate))

	return slices.ContainsFunc(values, func(value string) bool {
		return strings.ToLower(strings.TrimSpace(value)) == candidate
	})
}
