package database

import (
	"akerraps/tunectl/internal/cache"
	"akerraps/tunectl/internal/types"
	"context"
	"fmt"
)

func WriteSong (info *types.TrackInfo) (err error) {
	
	db,err:=cache.OpenDatabase()
	if err != nil {
		return  fmt.Errorf("cannot open database: %w", err)
	}

	_, err = db.ExecContext(context.Background(),
		`INSERT INTO songs (title, mbid) VALUES (?,?);`,
		info.Title, info.MBID)
	if err != nil {
		return fmt.Errorf("cannot insert song %q: %w", info.Title, err)
	}
	
return nil
}