package processor

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"
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

	//TODO: we may need to do some validation

	if len(DirFilePaths) == 0 {
		printError("no directories or files given")
		return
	}

	if Debug {
		printDebug(fmt.Sprintf("taking a look at the following paths: %v", DirFilePaths))
	}

	err := findMarkdownFiles(DirFilePaths)
	if err != nil {
		printError(err.Error())
	}
}

func processFlags() {
	if Debug {
		printDebug(fmt.Sprintf("Version: %s", Version))
		printDebug(fmt.Sprintf("LightMode: %v", LightMode))
		printDebug(fmt.Sprintf("Port: %s", Port))
	}
}

func findMarkdownFiles(paths []string) error {
	markdownFiles := make([]string, 0)
	for _, path := range paths {
		mdFiles, err := findFilesByType(path, ".md")
		if err != nil {
			if Debug {
				printDebug(fmt.Sprintf("unable to find markdown files at %s: %s", path, err.Error()))
			}
			printWarn(fmt.Sprintf("unable to find markdown files at %s", path))
		}

		markdownFiles = append(markdownFiles, mdFiles...)
	}

	fmt.Println("found md files: ")
	fmt.Println(markdownFiles)
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

//TODO: we could move all of this formatter logic over to a formatter package.

func printDebug(m string) {
	if Debug {
		fmt.Printf("DEBUG-%s: %s\n", getFormattedTime(), m)
	}
}

func printWarn(m string) {
	fmt.Printf("WARN-%s: %s\n", getFormattedTime(), m)
}

func printError(m string) {
	fmt.Printf("ERROR-%s: %s\n", getFormattedTime(), m)
}

// Get the time as standard UTC/Zulu format
func getFormattedTime() string {
	return time.Now().UTC().Format(time.RFC3339)
}
