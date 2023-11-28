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

// Grid is a grid of GitHub avatars.
//
// It expands and adds new rows when needed.
type Grid struct {
	// AvatarSize is the size of each avatar in the grid.
	AvatarSize int
	// Margin is the number of pixels between avatars.
	Margin int
	// Formatter is a function to call on avatar images to format them.
	Formatter avatar.Formatter
	// Image is the underlying image of the grid.
	image draw.Image
	// Row is the current row (0-indexed).
	row int
	// Col is the current column (0-indexed).
	col int
	// Cols is the number of columns in the grid.
	cols int
	// Rows is the number of rows in the grid.
	rows int
}

// WithSize updates the underlying image of the grid to fit the given number of
// avatars. This can help prevent frequent resizing of the underlying image.
func (g *Grid) WithSize(avatars int) {
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
	g.cols, g.rows = cols, rows
	g.image = g.newDst()
}

// Image returns the image of the grid.
func (g Grid) Image() image.Image {
	return g.image
}

// Cols returns the number of columns in the grid.
func (g Grid) Cols() int {
	return g.cols
}

// Rows returns the number of rows in the grid.
func (g Grid) Rows() int {
	return g.rows
}

// SetBounds changes the bounds of the underlying image.
func (g *Grid) setBounds(rows, cols int) {
	g.cols = cols
	g.rows = rows
	newImage := g.newDst()
	draw.Draw(newImage, newImage.Bounds(), g.image, image.Point{}, draw.Src)
	g.image = newImage
}

// NewDst creates a new destination image based on the grid's dimensions.
func (g Grid) newDst() draw.Image {
	width := g.cols*g.AvatarSize + (g.cols-1)*g.Margin
	height := g.rows*g.AvatarSize + (g.rows-1)*g.Margin
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// ColorModel returns the color model of the underlying image.
func (g Grid) ColorModel() color.Model {
	return g.image.ColorModel()
}

// Bounds returns the bounds of the underlying image.
func (g Grid) Bounds() image.Rectangle {
	return g.image.Bounds()
}

// At returns the color of the pixel at (x, y).
func (g Grid) At(x, y int) color.Color {
	return g.image.At(x, y)
}
