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

	browser := rod.New().MustConnect()
	defer browser.Close()

	page := browser.MustPage(url).MustSetViewport(1280, 720, 1, false)
	defer page.Close()

	err := proto.EmulationSetEmulatedMedia{
		Features: []*proto.EmulationMediaFeature{
			{Name: "prefers-color-scheme", Value: themeMode},
		},
	}.Call(page)

	if err != nil {
		return []byte{}, ErrScreenshotCapture
	}

	page.MustNavigate(url)

	page.MustWaitStable()
	buf := page.MustScreenshotFullPage("")

	if len(buf) == 0 {
		return []byte{}, ErrScreenshotCapture
	}

	return buf, nil
}
