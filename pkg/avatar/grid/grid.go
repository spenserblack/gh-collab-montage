// Package grid provides utilities for drawing a grid of GitHub avatars.
package grid

import (
	"image"
	"image/draw"

	"github.com/spenserblack/gh-collab-montage/pkg/avatar"
)

// PerRow is the number of avatars to draw per row.
const perRow = 10

// AvatarGrid is a grid of GitHub avatars.
//
// It expands and adds new rows when needed.
type AvatarGrid struct {
	image draw.Image
	row   int
	col   int
}

// New returns a new AvatarGrid.
func New() *AvatarGrid {
	return NewWithSize(0)
}

// NewWithSize returns a new AvatarGrid with the given size.
func NewWithSize(avatars int) *AvatarGrid {
	var cols, rows int
	if avatars == 0 {
		rows = 1
	} else {
		rows = avatars / perRow
	}
	// NOTE Always round the rows up.
	if avatars%perRow != 0 {
		rows++
	}
	// NOTE Max out at perRow columns.
	if avatars > perRow {
		cols = perRow
	} else {
		cols = avatars
	}
	return &AvatarGrid{
		image: image.NewRGBA(image.Rect(0, 0, cols*avatar.Width, rows*avatar.Height)),
	}
}

// Image returns the image of the grid.
func (g AvatarGrid) Image() image.Image {
	return g.image
}

// Cols returns the number of columns in the grid.
func (g AvatarGrid) Cols() int {
	return g.image.Bounds().Dx() / avatar.Width
}

// Rows returns the number of rows in the grid.
func (g AvatarGrid) Rows() int {
	return g.image.Bounds().Dy() / avatar.Height
}

// SetBounds changes the bounds of the underlying image.
func (g *AvatarGrid) setBounds(rows, cols int) {
	b := image.Rect(0, 0, cols*avatar.Width, rows*avatar.Height)
	newImage := image.NewRGBA(b)
	draw.Draw(newImage, b, g.image, image.Point{}, draw.Src)
	g.image = newImage
}
