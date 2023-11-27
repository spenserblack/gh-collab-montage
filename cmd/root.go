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
			avatars = append(avatars, a)
		}
		g := grid.NewWithSize(len(avatars), margin)
		for _, a := range avatars {
			g.AddAvatar(a)
		}

		m := g.Image()
		err = png.Encode(f, m)
		onError(err)
	},
}

// ImageStyleEnum is an enum for the different styles of images
type imageStyleEnum string

func (i imageStyleEnum) String() string {
	return string(i)
}

func (i *imageStyleEnum) Set(value string) error {
	switch value {
	case "":
		*i = "circle"
	case "circle", "square":
		*i = imageStyleEnum(value)
	default:
		return errors.New("invalid image style")
	}
	return nil
}

func (i imageStyleEnum) Type() string {
	return "circle | square"
}

var (
	margin     int
	imageStyle imageStyleEnum
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&margin, "margin", "m", 100, "Margin between avatars")
	rootCmd.PersistentFlags().VarP(&imageStyle, "style", "s", "Style of avatar (default circle)")
}
