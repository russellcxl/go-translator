package logger

import (
	"fmt"
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

type APILogger struct {
	logPath       string
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	fatalLogger   *log.Logger
}

func NewLogger(logPath string) *APILogger {
	return &APILogger{logPath: logPath}
}

func (l *APILogger) InitLogger() {

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(path.Join(l.logPath, "data.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l.infoLogger 	= log.New(file, "||--INFO-----|| ", log.Ldate|log.Ltime|log.Lshortfile)
	l.warningLogger = log.New(file, "||--WARNING--|| ", log.Ldate|log.Ltime|log.Lshortfile)
	l.errorLogger 	= log.New(file, "||--ERROR----|| ", log.Ldate|log.Ltime|log.Lshortfile)
	l.fatalLogger 	= log.New(file, "||--FATAL----|| ", log.Ldate|log.Ltime|log.Lshortfile)

}

func (l *APILogger) Info(args ...interface{}) {
	l.infoLogger.Output(2, fmt.Sprint(args...))
}

func (l *APILogger) Infof(template string, args ...interface{}) {
	l.infoLogger.Output(2, fmt.Sprintf(template, args...))
}

func (l *APILogger) Warn(args ...interface{}) {
	l.warningLogger.Output(2, fmt.Sprint(args...))
}

func (l *APILogger) Warnf(template string, args ...interface{}) {
	l.warningLogger.Output(2, fmt.Sprintf(template, args...))
}

func (l *APILogger) Error(args ...interface{}) {
	l.errorLogger.Output(2, fmt.Sprint(args...))
}

func (l *APILogger) Errorf(template string, args ...interface{}) {
	l.errorLogger.Output(2, fmt.Sprintf(template, args...))
}

func (l *APILogger) Fatal(args ...interface{}) {
	l.fatalLogger.Output(2, fmt.Sprint(args...))
	os.Exit(1)
}

func (l *APILogger) Fatalf(template string, args ...interface{}) {
	l.fatalLogger.Output(2, fmt.Sprintf(template, args...))
	os.Exit(1)
}
