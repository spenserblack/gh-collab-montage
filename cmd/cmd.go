// Package cmd contains the command line interface for the application.
package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	onError(rootCmd.Execute())
}

func onError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
