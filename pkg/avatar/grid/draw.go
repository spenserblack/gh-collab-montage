package grid

import (
	"image"

	"golang.org/x/image/draw"

	av "github.com/spenserblack/gh-collab-montage/pkg/avatar"
)

// AddAvatar adds an avatar's image to a grid.
//
// If needed, it expands the size of the underlying image.
func (g *Grid) AddAvatar(avatar image.Image) {
	resized := image.NewRGBA(image.Rect(0, 0, g.AvatarSize, g.AvatarSize))
	draw.ApproxBiLinear.Scale(resized, resized.Bounds(), avatar, avatar.Bounds(), draw.Src, nil)

	formatter := g.Formatter
	if formatter == nil {
		formatter = av.Noop
	}
	formatted := formatter(resized)
	// NOTE g.col and g.row are 0-indexed
	if g.row == 0 && g.Cols() <= g.col {
		g.setBounds(g.Rows(), g.Cols()+1)
	} else if g.col == 0 && g.Rows() <= g.row {
		g.setBounds(g.Rows()+1, g.Cols())
	}
	x := (g.col * g.AvatarSize) + (g.col * g.Margin)
	y := (g.row * g.AvatarSize) + (g.row * g.Margin)
	draw.Draw(
		g.image,
		image.Rect(x, y, x+g.AvatarSize, y+g.AvatarSize),
		formatted,
		image.Point{},
		draw.Src,
	)
	g.col++
	if g.col == perRow {
		g.col = 0
		g.row++
	}
}
