package loggers

import (
	"github.com/sirupsen/logrus"
	. "github.com/steallers/server-utility"
	"log"
)

type ErrorStructure struct {
	chain   string
	message error
}

type ErrorStructureInterface interface {
	Error() string
}

func manualPrinter(message string, action string, level string) {
	log.Printf("==================\n\tlevel:\t\t%s\n\tmessage: \t%s\n \taction:\t\t%s\n \tservice:\t%s\n\n======================================", level, message, action, ServiceName)
}

func PrintError(err error, action string) {
	if LoggerType == LoggerTypeJson {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.WithFields(logrus.Fields{
			"level":   "error",
			"action":  action,
			"service": ServiceName,
		}).Error(err)
	}

	if LoggerType == LoggerTypeText {
		logrus.SetFormatter(&logrus.TextFormatter{})
		manualPrinter(err.Error(), action, "error")
	}

	log.Print("invalid logger format")
}
