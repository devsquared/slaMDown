package main

import (
	"fmt"
	"os"

	"github.com/devsquared/slaMDown/processor"
	"github.com/spf13/cobra"
)

// TODO: let's utilize this main entry point as a base CLI. We then can create a server
//  package that will house the logic for a small local web server and we can create a
//  processor package that will hold the logic for the CLI commands.

func main() {
	fmt.Println("Welcome to slaMDown!")

	rootCommand := &cobra.Command{
		Use:     "slamdown [flags] [directory]",
		Short:   "slamdown [directory]",
		Long:    "slaMDown your markdown files from a directory to a parsed display locally",
		Version: processor.Version,
		Run: func(cmd *cobra.Command, args []string) {
			processor.DirFilePaths = args // pass directly
			processor.Process()
		},
	}

	flags := rootCommand.PersistentFlags()

	flags.BoolVarP(
		&processor.Debug,
		"debug",
		"d",
		false,
		"Use to switch to debug logging on.",
	)

	flags.BoolVarP(
		&processor.LightMode,
		"lightmode",
		"l",
		false,
		"Use to switch to light mode for the locally served files.",
	)

	flags.StringVarP(
		&processor.Port,
		"port",
		"p",
		"8080",
		"Define the port to be used when serving files locally.",
	)

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
