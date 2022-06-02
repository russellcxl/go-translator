package logger

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"log"
	"os"
	"testing"
	"time"
)

func Test_Logger(t *testing.T) {

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	//log.SetOutput(file)

	InfoLogger 		:= log.New(file, "||--INFO-----|| ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger 	:= log.New(file, "||--WARNING--|| ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger 	:= log.New(file, "||--ERROR----|| ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Println("Starting the application...")
	InfoLogger.Println("Something noteworthy happened")
	WarningLogger.Println("There is something you should know about")
	ErrorLogger.Println("Something went wrong")
}

func Test_Logrus(t *testing.T) {
	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
}

func Test_Zap(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "www.uber.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "www.uber.com")
}
