package fetcher

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

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
