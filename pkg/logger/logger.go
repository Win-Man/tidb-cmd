package logger

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	logLevel             = logrus.DebugLevel
	logPath              = "tidb-cmd.log"
	defaultLogTimeFormat = "2006/01/02 15:04:05.000"
)

// var (
// 	FileLog *log.Logger
// )

func InitLogger() {
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalln("Failed to open error log file:", err)
	}
	logrus.SetLevel(logLevel)
	logrus.AddHook(&contextHook{})
	logrus.SetFormatter(&contextFormatter{})
	logrus.SetOutput(file)

	// FileLog = log.New(file, "[", log.Ldate|log.Ltime|log.Lshortfile)
}

// func Debugf(msg string) {
// 	logrus.Debug(msg)
// }

// func Infof(msg string) {
// 	logrus.Info(msg)
// }

// func Warnf(msg string) {
// 	logrus.Warn(msg)
// }

// func Errorf(msg string) {
// 	logrus.Error(msg)
// }

// func Fatalf(msg string) {
// 	logrus.Fatal(msg)
// }

// func Panicf(msg string) {
// 	logrus.Panic(msg)
// }

type contextHook struct{}

func (hook *contextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 4)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !isSkippedPackageName(name) {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = filepath.Base(file)
			entry.Data["line"] = line
			break
		}
	}
	return nil
}

func (hook *contextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// isSKippedPackageName tests wether path name is on log library calling stack.
func isSkippedPackageName(name string) bool {
	return strings.Contains(name, "github.com/sirupsen/logrus") ||
		strings.Contains(name, "github.com/coreos/pkg/capnslog")
}

type contextFormatter struct{}

func (f *contextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	fmt.Fprintf(b, "[%s] ", entry.Time.Format(defaultLogTimeFormat))
	fmt.Fprintf(b, "[%s] ", entry.Level.String())
	if file, ok := entry.Data["file"]; ok {
		fmt.Fprintf(b, "[%s:%v] ", file, entry.Data["line"])
	}
	fmt.Fprintf(b, "[%s]", entry.Message)
	b.WriteByte('\n')

	return b.Bytes(), nil
}
