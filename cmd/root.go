package cmd

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"

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
		err := repoFlag.fillWithDefault()
		onError(err)
		f, err := os.Create("montage.png")
		onError(err)
		defer f.Close()
		client, err := api.DefaultRESTClient()
		onError(err)
		source := usersource.NewContributors(client, repoFlag.Owner, repoFlag.Name)
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
		switch avatarStyle.String() {
		case "circle":
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

		err = png.Encode(f, g)
		onError(err)
	},
}

// AvatarStyleEnum is an enum for the different styles of images
type avatarStyleEnum string

func (i avatarStyleEnum) String() string {
	if i == "" {
		return "circle"
	}
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

// Repo represents a GitHub repository
type repo struct {
	Owner string
	Name  string
}

func (r repo) String() string {
	if r.isEmpty() {
		return ""
	}
	return fmt.Sprintf("%s/%s", r.Owner, r.Name)
}

func (r *repo) Set(value string) error {
	split := strings.Split(value, "/")
	if len(split) != 2 {
		return errors.New("invalid repository format")
	}
	r.Owner, r.Name = split[0], split[1]
	return nil
}

func (repo) Type() string {
	return "OWNER/REPO"
}

// IsEmpty returns true if the repository hasn't been set.
func (r repo) isEmpty() bool {
	return r.Owner == "" && r.Name == ""
}

// FillWithDefault fills the repository with the current repository's info
// if the repository hasn't been set.
func (r *repo) fillWithDefault() error {
	if !r.isEmpty() {
		return nil
	}
	repository, err := repository.Current()
	if err != nil {
		return err
	}
	r.Owner, r.Name = repository.Owner, repository.Name
	return nil
}

var (
	margin      int
	avatarStyle avatarStyleEnum
	repoFlag    repo
)

func init() {
	rootCmd.PersistentFlags().VarP(&repoFlag, "repo", "R", "Specify another repository")
	rootCmd.PersistentFlags().IntVarP(&margin, "margin", "m", 100, "Margin between avatars")
	rootCmd.PersistentFlags().VarP(&avatarStyle, "style", "s", "Style of avatar")
}
