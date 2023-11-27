package cmd

import (
	"errors"
	"image"
	"image/png"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spenserblack/gh-collab-montage/pkg/avatar"
	"github.com/spenserblack/gh-collab-montage/pkg/avatar/grid"
	"github.com/spenserblack/gh-collab-montage/pkg/usersource"
	"github.com/spf13/cobra"
	"golang.org/x/image/draw"
)

var rootCmd = &cobra.Command{
	Use:   "gh-collab-montage",
	Short: "Combine your contributors avatars into a single image",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Create("montage.png")
		onError(err)
		defer f.Close()
		client, err := api.DefaultRESTClient()
		onError(err)
		repository, err := repository.Current()
		onError(err)
		source := usersource.NewContributors(client, repository.Owner, repository.Name)
		avatars := []image.Image{}
		for {
			user, stop, err := source.Next()
			onError(err)
			if stop {
				break
			}
			if user.Type != "User" {
				continue
			}
			a, err := avatar.Decode(user.AvatarURL)
			onError(err)
			// TODO Expose this to users
			resized := image.NewRGBA(image.Rect(0, 0, avatar.Width, avatar.Height))
			draw.NearestNeighbor.Scale(resized, resized.Bounds(), a, a.Bounds(), draw.Src, nil)
			avatars = append(avatars, resized)
		}

		var formatter avatar.Formatter
		switch avatarStyle {
		case "", "circle":
			formatter = avatar.Circlify
		case "square":
			formatter = avatar.Noop
		default:
			panic("unreachable: invalid avatar style")
		}
		g := grid.NewWithSize(len(avatars), margin, formatter)
		for _, a := range avatars {
			g.AddAvatar(a)
		}

		m := g.Image()
		err = png.Encode(f, m)
		onError(err)
	},
}

// AvatarStyleEnum is an enum for the different styles of images
type avatarStyleEnum string

func (i avatarStyleEnum) String() string {
	return string(i)
}

func (i *avatarStyleEnum) Set(value string) error {
	switch value {
	case "circle", "square":
		*i = avatarStyleEnum(value)
	default:
		return errors.New("invalid image style")
	}
	return nil
}

func (i avatarStyleEnum) Type() string {
	return "circle | square"
}

var (
	margin      int
	avatarStyle avatarStyleEnum
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&margin, "margin", "m", 100, "Margin between avatars")
	rootCmd.PersistentFlags().VarP(&avatarStyle, "style", "s", "Style of avatar (default circle)")
}
