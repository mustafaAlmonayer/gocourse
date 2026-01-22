package intermediate

import (
	"io"
	"log"
	"os"
)

// Production ready package is logurs and zap

var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	logFile     *os.File
)

func init() {
	var err error
	logFile, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("Error creating logging file: %v", err)
	}

	out := io.MultiWriter(os.Stdout, logFile)
	flags := log.Ldate | log.Ltime | log.Lshortfile

	InfoLogger = log.New(out, "INFO: ", flags)
	WarnLogger = log.New(out, "WARN: ", flags)
	ErrorLogger = log.New(out, "ERROR: ", flags)
}

func logging() {
	defer logFile.Close()
	InfoLogger.Println("This is info")
	WarnLogger.Println("This is warn")
	ErrorLogger.Println("This is error")
}
