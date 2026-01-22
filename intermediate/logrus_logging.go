package intermediate

import "github.com/sirupsen/logrus"

var (
	LogrusLogging = logrus.New()
)

func logrusLogging() {

	// Set log level
	LogrusLogging.SetLevel(logrus.InfoLevel)

	// Format
	LogrusLogging.SetFormatter(&logrus.JSONFormatter{})

	// Exps
	LogrusLogging.Info("This is a info message")
	LogrusLogging.Warn("This is a warn message")
	LogrusLogging.Error("This is a error message")

	LogrusLogging.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("User Logged in.")

}
