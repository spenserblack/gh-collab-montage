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
	// Row is the current row (0-indexed).
	row int
	// Col is the current column (0-indexed).
	col int
	// Margin is the number of pixels between avatars.
	margin int
	// Cols is the number of columns in the grid.
	cols int
	// Rows is the number of rows in the grid.
	rows int
}

// New returns a new AvatarGrid.
func New(margin int) *AvatarGrid {
	return NewWithSize(0, margin)
}

// NewWithSize returns a new AvatarGrid with the given size.
func NewWithSize(avatars int, margin int) *AvatarGrid {
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
	width := cols*avatar.Width + (cols-1)*margin
	height := rows*avatar.Height + (rows-1)*margin
	return &AvatarGrid{
		image:  image.NewRGBA(image.Rect(0, 0, width, height)),
		margin: margin,
		cols:   cols,
		rows:   rows,
	}
}

// Image returns the image of the grid.
func (g AvatarGrid) Image() image.Image {
	return g.image
}

// Cols returns the number of columns in the grid.
func (g AvatarGrid) Cols() int {
	return g.cols
}

// Rows returns the number of rows in the grid.
func (g AvatarGrid) Rows() int {
	return g.rows
}

// SetBounds changes the bounds of the underlying image.
func (g *AvatarGrid) setBounds(rows, cols int) {
	g.cols = cols
	g.rows = rows
	width := (cols * avatar.Width) + ((cols - 1) * g.margin)
	height := (rows * avatar.Height) + ((rows - 1) * g.margin)
	b := image.Rect(0, 0, width, height)
	newImage := image.NewRGBA(b)
	draw.Draw(newImage, b, g.image, image.Point{}, draw.Src)
	g.image = newImage
}
