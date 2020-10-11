package console

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// LogMux multiplexes logs into several handlers. It prints logs to stderr by
// default.
type LogMux struct {
	logger *log.Logger

	mu       sync.Mutex
	handlers []func(Level, string, ...interface{}) bool
}

// Level defines log level.
type Level int

// Log levels
const (
	LLInfo Level = iota
	LLWarning
	LLError
	LLFatal
)

var logMux = &LogMux{
	logger: log.New(os.Stderr, "", log.LstdFlags),
}

func (m *LogMux) addHandler(h func(Level, string, ...interface{}) bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.handlers = append(m.handlers, h)
}

func (m *LogMux) handleAll(l Level, format string, v ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, handler := range m.handlers {
		if res := handler(l, format, v...); !res {
			m.logger.Printf(SLogf(LLError, "failed to handle log"))
		}
	}
}

// Info leaves log by log level.
func Info(format string, v ...interface{}) {
	Printf(SLogf(LLInfo, format, v...))
	logMux.handleAll(LLInfo, format, v...)
}

// Warning leaves log by log level.
func Warning(format string, v ...interface{}) {
	Printf(SLogf(LLWarning, format, v...))
	logMux.handleAll(LLWarning, format, v...)
}

// Error leaves log by log level.
func Error(format string, v ...interface{}) {
	Printf(SLogf(LLError, format, v...))
	logMux.handleAll(LLError, format, v...)
}

// Fatal leaves log by log level and exits the process.
func Fatal(format string, v ...interface{}) {
	Printf(SLogf(LLFatal, format, v...))
	logMux.handleAll(LLFatal, format, v...)
	os.Exit(1)
}

// AddLogHandler adds addtional log handling functions. Added handlers are applied
// to every logging functions like Info, Warning, Error, Fatal.
func AddLogHandler(handler func(Level, string, ...interface{}) bool) {
	logMux.addHandler(handler)
}

// Printf just prints input into stderr without any log level.
func Printf(format string, v ...interface{}) {
	logMux.logger.Printf(format, v...)
}

// SLogf formats logs by Level and return the result string.
// For general purpose, you should use Info, Warning, Error, Fatal function.
func SLogf(l Level, format string, v ...interface{}) string {
	var formatted string

	switch l {
	case LLInfo:
		formatted = fmt.Sprintf("[INFO] "+format, v...)
	case LLWarning:
		formatted = fmt.Sprintf("[WARNING] "+format, v...)
	case LLError:
		formatted = fmt.Sprintf("[ERROR] "+format, v...)
	case LLFatal:
		formatted = fmt.Sprintf("[FATAL] "+format, v...)
	}
	return formatted
}
