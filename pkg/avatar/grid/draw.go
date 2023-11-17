package grid

import "image"

// AddAvatar adds an avatar's image to a grid.
//
// If needed, it expands the size of the underlying image.
func (g *AvatarGrid) AddAvatar(avatar image.Image) {
	// TODO Assert that avatars are the appropriate size?
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
