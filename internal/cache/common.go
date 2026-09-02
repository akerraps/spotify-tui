package cache

import (
	"fmt"
	"os"
	"os/user"
)

func getCachePath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("cannot get current user: %w", err)
	}
	cachePath := fmt.Sprintf("/home/%s/.cache/tunectl/", currentUser.Username)
	return cachePath, nil
}

func createCacheDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func InitCache()(bin string, err error ){
	cachePath, err := getCachePath()
	if err != nil {
		return "", fmt.Errorf("cannot get cache path: %w", err)
	}
	
	if err := createCacheDir(cachePath); err != nil {
		return "", fmt.Errorf("cannot create cache dir: %w", err)
	}

	db, err := OpenDatabase()
	if err != nil {
		return "", fmt.Errorf("cannot open database: %w", err)
	}
	
	if err := initDatabase(db); err != nil {
		return "", fmt.Errorf("cannot init database: %w", err)
	}
	
	bin, err = getYtDlp(cachePath)
	if err != nil {
    return "", fmt.Errorf("cannot initialize yt-dlp: %w", err)
	}
	
	defer db.Close()
	
	return bin, nil
}