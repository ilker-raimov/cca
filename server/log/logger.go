package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Formatter struct{}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s [%s]: %s\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level, entry.Message)), nil
}

func Init() {
	logrus.SetFormatter(&Formatter{})
}

func Info(log string) {
	logrus.Info(log)
}

func InfoF(format string, parameters ...any) {
	log := fmt.Sprintf(format, parameters...)

	Info(log)
}
