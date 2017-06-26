package main

import (
	"github.com/Sirupsen/logrus"
)

func setupLogging() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}
