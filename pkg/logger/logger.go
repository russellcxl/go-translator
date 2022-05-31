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
}

type logger struct {
	logPath string
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

func NewLogger(logPath string) *logger {
	return &logger{logPath: logPath}
}

func (l *logger) InitLogger() {

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(path.Join(l.logPath, "data.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "||--INFO-----|| ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "||--WARNING--|| ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "||--ERROR----|| ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(file, "||--FATAL----|| ", log.Ldate|log.Ltime|log.Lshortfile)

}

func (l *logger) Info(args ...interface{}) {
	InfoLogger.Println(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	InfoLogger.Printf(template, args...)
}

func (l *logger) Warn(args ...interface{}) {
	WarningLogger.Println(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	WarningLogger.Printf(template, args...)
}

func (l *logger) Error(args ...interface{}) {
	ErrorLogger.Println(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	ErrorLogger.Printf(template, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	ErrorLogger.Fatalln(args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	ErrorLogger.Fatalf(template, args...)
}
