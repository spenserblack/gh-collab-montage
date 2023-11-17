// Package grid provides utilities for drawing a grid of GitHub avatars.
package grid

import (
	"image"
	"image/draw"
)

const (
	// AvatarWidth is the width of an avatar in pixels.
	AvatarWidth = 500
	// AvatarHeight is the height of an avatar in pixels.
	AvatarHeight = 500
	// PerRow is the number of avatars to draw per row.
	perRow = 10
)

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
		image: image.NewRGBA(image.Rect(0, 0, cols*AvatarWidth, rows*AvatarHeight)),
	}
}

// Image returns the image of the grid.
func (g AvatarGrid) Image() image.Image {
	return g.image
}

// Cols returns the number of columns in the grid.
func (g AvatarGrid) Cols() int {
	return g.image.Bounds().Dx() / AvatarWidth
}

// Rows returns the number of rows in the grid.
func (g AvatarGrid) Rows() int {
	return g.image.Bounds().Dy() / AvatarHeight
}

// IncAvatar increments to the next avatar.
//
// It expands the grid and adds new columns and rows when needed.
func (g *AvatarGrid) incAvatar() {
	// NOTE g.col and g.row are 0-indexed
	if g.row == 0 && g.Cols() <= g.col {
		g.setBounds(g.Rows(), g.Cols()+1)
	} else if g.col == 0 && g.Rows() <= g.row {
		g.setBounds(g.Rows()+1, g.Cols())
	}
	g.col++
	if g.col == perRow {
		g.col = 0
		g.row++
	}
}

// SetBounds changes the bounds of the underlying image.
func (g *AvatarGrid) setBounds(rows, cols int) {
	b := image.Rect(0, 0, cols*AvatarWidth, rows*AvatarHeight)
	newImage := image.NewRGBA(b)
	draw.Draw(newImage, b, g.image, image.Point{}, draw.Src)
	g.image = newImage
}
