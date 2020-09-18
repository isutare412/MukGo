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

func (l *logger) toConsole(time time.Time, msg string) {
	console.InfoTag(l.tag, "%s", msg)
}

func (l *logger) log(time time.Time, msg string) error {
	l.toConsole(time, msg)
	err := l.toFile(time, msg)
	return err
}
