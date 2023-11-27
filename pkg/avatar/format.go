package avatar

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Formatter is a function to call on avatar images to format them.
type Formatter = func(image.Image) image.Image

var transparent = color.Alpha{0}

// Circlify is a utility function to convert an avatar to a circle.
func Circlify(avatar image.Image) image.Image {
	circle := image.NewRGBA(avatar.Bounds())
	draw.Draw(circle, avatar.Bounds(), avatar, image.Point{}, draw.Src)

	origin := struct {
		x float64
		y float64
	}{
		x: float64(avatar.Bounds().Dx() / 2),
		y: float64(avatar.Bounds().Dy() / 2),
	}
	radius := math.Min(origin.x, origin.y)
	for x := 0; x < avatar.Bounds().Dx(); x++ {
		for y := 0; y < avatar.Bounds().Dy(); y++ {
			distance := math.Sqrt(math.Pow(float64(x)-origin.x, 2) + math.Pow(float64(y)-origin.y, 2))
			if distance > radius {
				circle.Set(x, y, transparent)
			}
		}
	}
	return circle
}

// Noop is a utility function to do nothing to an avatar. This exists to have the same
// signature as other formatting functions.
func Noop(avatar image.Image) image.Image {
	return avatar
}
