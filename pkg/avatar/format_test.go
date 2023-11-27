package avatar

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func TestCirclify(t *testing.T) {
	// NOTE Creating an image with a predictable color
	baseImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	baseColor := color.Gray{0xFF}
	transparent := color.Alpha{0x00}
	draw.Draw(baseImg, baseImg.Bounds(), image.NewUniform(baseColor), image.Point{}, draw.Src)

	circle := Circlify(baseImg)
	if circle.Bounds() != baseImg.Bounds() {
		t.Fatalf("circle.Bounds() = %v, want %v", circle.Bounds(), baseImg.Bounds())
	}

	tests := []struct {
		x, y int
		want color.Color
	}{
		{0, 0, transparent},
		{50, 50, baseColor},
		{99, 99, transparent},
		
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("(%d, %d)", tt.x, tt.y), func(t *testing.T) {
			got := circle.At(tt.x, tt.y)
			AssertColorsEqual(t, got, tt.want)
		})
	}
}

// AssertColorsEqual asserts that two colors are equal.
func AssertColorsEqual(t *testing.T, actual, want color.Color) {
	t.Helper()
	r, g, b, a := actual.RGBA()
	wr, wg, wb, wa := want.RGBA()
	actualRGBA := [4]uint32{r, g, b, a}
	wantRGBA := [4]uint32{wr, wg, wb, wa}

	if actualRGBA != wantRGBA {
		t.Errorf("RGBA = %v, want %v", actualRGBA, wantRGBA)
	}
}
