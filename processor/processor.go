package processor

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/devsquared/slaMDown/formatter"
	"github.com/devsquared/slaMDown/server"
	"github.com/devsquared/slaMDown/util"
)

// Version is the version of the application. Formatted as semver - https://semver.org/
var Version = "0.1.0"

// DirFilePaths is set from args to the application; not flag. These are a slice of strings and will be validated.
var DirFilePaths []string

// Debug is a flag used to turn on verbose logging.
var Debug = false

// LightMode is a variable indicating if the user wants to use the application in light or dark mode.
var LightMode = false

// Port is the port used when serving locally to the web.
var Port = "8080"

// Process is the main entry point for our application.
func Process() {
	processFlags() // start with this to properly debug out context of the application

	if len(DirFilePaths) == 0 {
		formatter.PrintError("no directories or files given")
		return
	}

	if Debug {
		formatter.PrintDebug(fmt.Sprintf("taking a look at the following paths: %v", DirFilePaths))
	}

	err := findMarkdownFiles(DirFilePaths)
	if err != nil {
		formatter.PrintError(err.Error())
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, util.DebugContextKey, Debug)
	ctx = context.WithValue(ctx, util.PortContextKey, Port)
	server.NewServer(ctx)
}

func processFlags() {
	if Debug {
		formatter.PrintDebug(fmt.Sprintf("Version: %s", Version))
		formatter.PrintDebug(fmt.Sprintf("LightMode: %v", LightMode))
		formatter.PrintDebug(fmt.Sprintf("Port: %s", Port))
	}
}

func findMarkdownFiles(paths []string) error {
	markdownFiles := make([]string, 0)
	for _, path := range paths {
		mdFiles, err := findFilesByType(path, ".md")
		if err != nil {
			if Debug {
				formatter.PrintDebug(fmt.Sprintf("unable to find markdown files at %s: %s", path, err.Error()))
			}
			formatter.PrintWarn(fmt.Sprintf("unable to find markdown files at %s", path))
		}

		markdownFiles = append(markdownFiles, mdFiles...)
	}

	if Debug {
		formatter.PrintDebug("found md files: ")
		formatter.PrintDebug(fmt.Sprintf("%s", markdownFiles))
	}

	return nil
}

func findFilesByType(path, fType string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == fType {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
