package main

import (
	"fmt"
	"os"

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
		Version: "1", // TODO: make a processor for this?
		Run: func(cmd *cobra.Command, args []string) {
			//TODO: here we will call any setup we want from the processor package

			//TODO: and then one final processor.Process()
			fmt.Println("processing the commands")
		},
	}

	//TODO: utilize this space to add any flags that we want. As of now, those are:
	// - port
	// - light (for lightmode)

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
