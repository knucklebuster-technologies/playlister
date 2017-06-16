package loggy

import (
	"io"
	"log"
	"os"
)

// Logger simple struct to log events
type Logger struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

// New sets up a Logger for use through out application
func New(infoHandle io.Writer, warnHandle io.Writer, errorHandle io.Writer) *Logger {
	return &Logger{
		log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(warnHandle, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Fatal calls Error.Println followed by os.Exit(1)
func (l *Logger) Fatal(i interface{}) {
	l.Error.Print(i)
	os.Exit(1)
}
