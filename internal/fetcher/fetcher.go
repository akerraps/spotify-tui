package fetcher

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-getter"
	"github.com/kkdai/youtube/v2"
)

func TcheckDependencies() error {
	currentUser, err1 := user.Current()
	if err1 != nil {
		return err1
	}

	binary := "yt-dlp"
	cachePath := fmt.Sprintf("/home/%s/.cache/tunectl/", currentUser.Username)

	client := &getter.Client{
		Dst:   cachePath + binary,
		Src:   "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp",
		Mode:  getter.ClientModeFile,
		Umask: 0o077,
	}

	if err2 := client.Get(); err2 != nil {
		return err2
	}

	path := filepath.Join(cachePath, "yt-dlp")
	if err3 := os.Chmod(path, 0o700); err3 != nil {
		return fmt.Errorf("cannot make yt-dlp executable: %w", err3)
	}

	return nil
}

func FetchAudio(videoID string) error {
	client := youtube.Client{}

	// Fetch video metadata (no download yet)
	video, err := client.GetVideo(videoID)
	if err != nil {
		return err
	}

	// Filter audio-only formats
	var audioFormats youtube.FormatList

	for i := range video.Formats {
		f := video.Formats[i]

		// Audio-only formats:
		// - Have audio channels
		// - Do NOT have QualityLabel (video formats always have it)
		if f.AudioChannels > 0 && f.QualityLabel == "" {
			audioFormats = append(audioFormats, f)
		}
	}

	if len(audioFormats) == 0 {
		return errors.New("no audio-only format found")
	}

	// Pick best quality audio (highest bitrate)
	audioFormats.Sort()
	best := audioFormats[0]

	// Detect correct file extension from mime type
	ext := ".m4a"
	if strings.Contains(best.MimeType, "webm") {
		ext = ".webm"
	}

	// Open audio stream
	stream, _, err := client.GetStream(video, &best)
	if err != nil {
		return err
	}
	defer stream.Close()

	file, err := os.Create("audio" + ext)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	return err
}
