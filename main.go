package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spenserblack/gh-collab-montage/pkg/avatar"
	"github.com/spenserblack/gh-collab-montage/pkg/avatar/grid"
	"github.com/spenserblack/gh-collab-montage/pkg/usersource"
)

func main() {
	f, err := os.Create("montage.png")
	defer f.Close()
	onError(err)
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
	g := grid.NewWithSize(len(avatars))
	for _, a := range avatars {
		g.AddAvatar(a)
	}

	m := g.Image()
	err = png.Encode(f, m)
	onError(err)
}

func onError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
