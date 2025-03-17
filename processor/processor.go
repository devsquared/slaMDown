package processor

import (
	"fmt"
	"time"
)

// Version is the version of the application.
var Version = "1.0.0"

// DirFilePaths is set from args to the application; not flag. These are a slice of strings and will be validated.
var DirFilePaths []string

// Debug is a flag used to turn on verbose logging.
var Debug = false

// LightMode is a variable indicating if the user wants to use the application in light or dark mode.
var LightMode = false

// Port is the port used when serving locally to the web.
var Port = "8080"

func Process() {
	processFlags() // start with this to properly debug out context of the application

	if len(DirFilePaths) == 0 {
		printError("no directories or files given")
	}

	if Debug {
		printDebug(fmt.Sprintf("taking a look at the following paths: %v", DirFilePaths))
	}
}

func processFlags() {
	if Debug {
		printDebug(fmt.Sprintf("Version: %s", Version))
		printDebug(fmt.Sprintf("LightMode: %v", LightMode))
		printDebug(fmt.Sprintf("Port: %s", Port))
	}
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
