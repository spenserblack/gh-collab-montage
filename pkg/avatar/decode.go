package avatar

import (
	"image"
	_ "image/png"
	"net/http"
)

// DecodeAvatar decodes a GitHub avatar from a URL.
func DecodeAvatar(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	m, _, err := image.Decode(resp.Body)
	return m, err
}
