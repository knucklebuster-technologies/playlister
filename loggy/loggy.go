package loggy

import (
	"log"
	"os"
)

// Info logs information events
var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

// Warn logs warning events
var Warn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

// Error logs error events
var Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

// Fatal same as Error call followed by os.Exit(1)
func Fatal(i interface{}) {
	Error.Print(i)
	os.Exit(1)
}
