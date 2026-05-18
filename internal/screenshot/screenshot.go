package screenshot

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func Screenshot(url string, theme ...string) ([]byte, error) {
	themeMode := "light"
	if len(theme) != 0 && theme[0] == "dark" {
		themeMode = "dark"
	}

	if url == "" {
		return []byte{}, ErrMissingPortfolioURL
	}

	page := rod.New().MustConnect().MustPage(url).MustSetViewport(1280, 720, 1, false)

	err := proto.EmulationSetEmulatedMedia{
		Features: []*proto.EmulationMediaFeature{
			{Name: "prefers-color-scheme", Value: themeMode},
		},
	}.Call(page)

	if err != nil {
		return []byte{}, ErrScreenshotCapture
	}

	buf := page.MustWaitStable().MustScreenshotFullPage("")

	if len(buf) == 0 {
		return []byte{}, ErrScreenshotCapture
	}

	return buf, nil
}
