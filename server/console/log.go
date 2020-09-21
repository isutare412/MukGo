package console

import (
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", log.LstdFlags)

// Info leaves log by log level.
func Info(format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("[INFO] %s", format), v...)
}

// InfoTag leaves log by log level with custom tag.
func InfoTag(tag string, format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("[%s] %s", tag, format), v...)
}

// Warning leaves log by log level.
func Warning(format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("[WARNING] %s", format), v...)
}

// Error leaves log by log level.
func Error(format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("[ERROR] %s", format), v...)
}

// Fatal leaves log by log level and exits the process.
func Fatal(format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("[FATAL] %s", format), v...)
	os.Exit(1)
}
