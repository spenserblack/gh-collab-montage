package grid

import (
	"fmt"
	"image"
	"testing"

	av "github.com/spenserblack/gh-collab-montage/pkg/avatar"
)

func TestWithSize(t *testing.T) {
	tests := []struct {
		name    string
		avatars int
		cols    int
		rows    int
	}{
		{
			name:    "zero avatars",
			avatars: 0,
			cols:    0,
			rows:    1,
		},
		{
			name:    "one avatar",
			avatars: 1,
			cols:    1,
			rows:    1,
		},
		{
			name:    "ten avatars",
			avatars: 10,
			cols:    10,
			rows:    1,
		},
		{
			name:    "eleven avatars",
			avatars: 11,
			cols:    10,
			rows:    2,
		},
		{
			name:    "twenty avatars",
			avatars: 20,
			cols:    10,
			rows:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				AvatarSize: 400,
				Margin:     100,
				Formatter:  av.Noop,
			}
			g.WithSize(tt.avatars)
			if g.Cols() != tt.cols {
				t.Errorf("g.Cols() = %d, want %d", g.Cols(), tt.cols)
			}
			if g.Rows() != tt.rows {
				t.Errorf("g.Rows() = %d, want %d", g.Rows(), tt.rows)
			}
		})
	}
}

// Tests that size is adjusted when needed.
func TestGrid_AddAvatar(t *testing.T) {
	// TODO Test grid's pixels by drawing images with known colors
	avatar := image.NewAlpha(image.Rect(0, 0, 500, 500))
	tests := []struct {
		size int
		n    int
		cols int
		rows int
		col  int
		row  int
	}{
		{
			size: 0,
			n:    1,
			cols: 1,
			rows: 1,
			col:  1,
			row:  0,
		},
		{
			size: 1,
			n:    1,
			cols: 1,
			rows: 1,
			col:  2,
			row:  0,
		},
		{
			size: 1,
			n:    2,
			cols: 2,
			rows: 1,
			col:  3,
			row:  0,
		},
		{
			size: 10,
			n:    10,
			cols: 10,
			rows: 1,
			col:  0,
			row:  1,
		},
		{
			size: 0,
			n:    11,
			cols: 10,
			rows: 2,
			col:  1,
			row:  1,
		},
		{
			size: 10,
			n:    11,
			cols: 10,
			rows: 2,
			col:  1,
			row:  1,
		},
		{
			size: 11,
			n:    12,
			cols: 10,
			rows: 2,
			col:  2,
			row:  1,
		},
		{
			size: 0,
			n:    20,
			cols: 10,
			rows: 2,
			col:  0,
			row:  2,
		},
		{
			size: 20,
			n:    20,
			cols: 10,
			rows: 2,
			col:  0,
			row:  2,
		},
		{
			size: 0,
			n:    21,
			cols: 10,
			rows: 3,
			col:  1,
			row:  2,
		},
		{
			size: 20,
			n:    21,
			cols: 10,
			rows: 3,
			col:  1,
			row:  2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d avatars added to %d-avatar grid", tt.n, tt.size), func(t *testing.T) {
			g := &Grid{
				AvatarSize: 400,
				Margin:     100,
				Formatter:  av.Noop,
			}
			g.WithSize(tt.size)
			for i := 0; i < tt.n; i++ {
				g.AddAvatar(avatar)
			}
			if g.Cols() != tt.cols {
				t.Errorf("tt.g.Cols() = %d, want %d", g.Cols(), tt.cols)
			}
			if g.Rows() != tt.rows {
				t.Errorf("tt.g.Rows() = %d, want %d", g.Rows(), tt.rows)
			}
		})
	}
}

// Tests that the Noop formatter is used when no formatter is provided.
func TestGrid_AddAvatar_nil_formatter(t *testing.T) {
	avatar := image.NewAlpha(image.Rect(0, 0, 500, 500))
	g := &Grid{
		AvatarSize: 400,
		Margin:     100,
	}
	g.WithSize(1)
	g.AddAvatar(avatar)
	// NOTE Basically if we didn't panic from a nil pointer dereference, we're good
	// TODO Test grid's pixels by drawing images with known colors
}
