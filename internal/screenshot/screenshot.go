package screenshot

import (
	"errors"

	"github.com/go-rod/rod"
)

func Screenshot(url string) ([]byte, error) {
	page := rod.New().MustConnect()

	buf := page.MustPage(url).MustWaitStable().MustSetViewport(1280, 720, 1, false).MustScreenshotFullPage("")
	if len(buf) == 0 {
		return []byte{}, errors.New("Failed to capture portfolio screenshot.")
	}

	return buf, nil
}
