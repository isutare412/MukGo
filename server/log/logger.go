package log

import (
	"fmt"
	"time"

	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/dailyrotate"
)

type logger struct {
	sender string
	tag    string

	dir  string
	file *dailyrotate.File
}

func (l *logger) toFile(time time.Time, msg string) error {
	// add newline if not exists
	format := "%s %s"
	if len(msg) == 0 || msg[len(msg)-1] != '\n' {
		format = "%s %s\n"
	}

	line := fmt.Sprintf(format, time.Format("[15:04:05]"), msg)

	_, err := l.file.Write([]byte(line))
	return err
}

func (l *logger) toConsole(level console.Level, msg string) {
	logBody := console.SLogf(level, msg)
	console.Printf(fmt.Sprintf("[%s]%s", l.tag, logBody))
}

func (l *logger) log(time time.Time, level console.Level, msg string) error {
	l.toConsole(level, msg)
	err := l.toFile(time, msg)
	return err
}
