package logger

import (
	"log"
	"os"
	"path"
)

// Logger methods interface
type Logger interface {
	InitLogger()
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type logger struct {
	logPath       string
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
}

func NewLogger(logPath string) *logger {
	return &logger{logPath: logPath}
}

func (l *logger) InitLogger() {

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(path.Join(l.logPath, "data.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l.InfoLogger = log.New(file, "||--INFO-----|| ", log.Ldate|log.Ltime|log.Lshortfile)
	l.WarningLogger = log.New(file, "||--WARNING--|| ", log.Ldate|log.Ltime|log.Lshortfile)
	l.ErrorLogger = log.New(file, "||--ERROR----|| ", log.Ldate|log.Ltime|log.Lshortfile)
	l.FatalLogger = log.New(file, "||--FATAL----|| ", log.Ldate|log.Ltime|log.Lshortfile)

}

func (l *logger) Info(args ...interface{}) {
	l.InfoLogger.Println(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.InfoLogger.Printf(template, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.WarningLogger.Println(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.WarningLogger.Printf(template, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.ErrorLogger.Println(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.ErrorLogger.Printf(template, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.FatalLogger.Fatalln(args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.FatalLogger.Fatalf(template, args...)
}
