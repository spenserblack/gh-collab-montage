package avatar

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

// Decode decodes a GitHub avatar from a URL.
func Decode(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	m, _, err := image.Decode(resp.Body)
	return m, err
}
