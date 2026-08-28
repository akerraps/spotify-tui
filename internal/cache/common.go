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