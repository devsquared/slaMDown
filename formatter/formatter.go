package formatter

import (
	"fmt"
	"time"
)

func PrintDebug(m string) {
	fmt.Printf("DEBUG-%s: %s\n", getFormattedTime(), m)
}

func PrintWarn(m string) {
	fmt.Printf("WARN-%s: %s\n", getFormattedTime(), m)
}

func PrintError(m string) {
	fmt.Printf("ERROR-%s: %s\n", getFormattedTime(), m)
}

// Get the time as standard UTC/Zulu format
func getFormattedTime() string {
	return time.Now().UTC().Format(time.RFC3339)
}
