// Package grid provides utilities for drawing a grid of GitHub avatars.
package grid

import (
	"image"
	"image/color"

	"golang.org/x/image/draw"

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
	// Formatter is a function to call on avatar images to format them.
	formatter avatar.Formatter
}

// New returns a new AvatarGrid.
func New(margin int, formatter avatar.Formatter) *AvatarGrid {
	return NewWithSize(0, margin, formatter)
}

// NewWithSize returns a new AvatarGrid with the given size.
func NewWithSize(avatars int, margin int, formatter avatar.Formatter) *AvatarGrid {
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
	g := &AvatarGrid{
		margin:    margin,
		cols:      cols,
		rows:      rows,
		formatter: formatter,
	}
	g.image = g.newDst()
	return g
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
	newImage := g.newDst()
	draw.Draw(newImage, newImage.Bounds(), g.image, image.Point{}, draw.Src)
	g.image = newImage
}

// NewDst creates a new destination image based on the grid's dimensions.
func (g AvatarGrid) newDst() draw.Image {
	width := g.cols*avatar.Width + (g.cols-1)*g.margin
	height := g.rows*avatar.Height + (g.rows-1)*g.margin
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// ColorModel returns the color model of the underlying image.
func (g AvatarGrid) ColorModel() color.Model {
	return g.image.ColorModel()
}

// Bounds returns the bounds of the underlying image.
func (g AvatarGrid) Bounds() image.Rectangle {
	return g.image.Bounds()
}

// At returns the color of the pixel at (x, y).
func (g AvatarGrid) At(x, y int) color.Color {
	return g.image.At(x, y)
}
