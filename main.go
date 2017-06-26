package main

import (
	"time"

	"github.com/Sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting query at time: %s", time.Now())
}
