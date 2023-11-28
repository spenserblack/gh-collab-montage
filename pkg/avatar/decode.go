package avatar

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

// DecodeFromURL decodes an image GitHub avatar from a URL.
//
// This can be used to get an avatar from GitHub.
func DecodeFromURL(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	m, _, err := image.Decode(resp.Body)
	return m, err
}
