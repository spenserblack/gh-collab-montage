package main

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spenserblack/gh-collab-montage/pkg/usersource"
)

func main() {
	client, err := api.DefaultRESTClient()
	onError(err)
	repository, err := repository.Current()
	onError(err)
	source := usersource.NewContributors(client, repository.Owner, repository.Name)
	for {
		user, stop, err := source.Next()
		onError(err)
		if stop {
			break
		}
		fmt.Println(user)
	}
}

func onError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
