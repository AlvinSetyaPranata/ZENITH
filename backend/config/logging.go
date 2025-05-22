package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type LEVEL struct {
	info     *log.Logger
	debug    *log.Logger
	warning  *log.Logger
	critical *log.Logger
}

var Log *LEVEL

func init() {
	Log = &LEVEL{
		info:     log.New(os.Stdout, "", 0),
		debug:    log.New(os.Stdout, "", 0),
		warning:  log.New(os.Stdout, "", 0),
		critical: log.New(os.Stdout, "", 0),
	}
}

func format(level string, messege string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s][%s]: %s", t, level, messege)
}

func (l *LEVEL) Info(msg ...string) {
	l.info.Println(format("INFO", strings.Join(msg, "")))
}

func (l *LEVEL) Debug(msg ...string) {
	l.debug.Println(format("debug", strings.Join(msg, "")))
}
func (l *LEVEL) Warning(msg ...string) {
	l.warning.Println(format("INFO", strings.Join(msg, "")))
}

func (l *LEVEL) Critical(msg ...string) {
	l.critical.Println(format("INFO", strings.Join(msg, "")))
}
