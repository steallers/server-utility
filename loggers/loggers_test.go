package loggers

import (
	"github.com/sirupsen/logrus"
	"testing"
)


func TestLogging_Test1(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
