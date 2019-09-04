package utils

import (
	"github.com/stock-details-api/internal/constants"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var once sync.Once
var instance logrus.Logger

func GetLogger() logrus.Logger {
	once.Do(func() {
		instance = initializeLogger()
	})
	return instance
}

func initializeLogger() logrus.Logger {
	log := logrus.New()
	f, err := os.OpenFile(constants.LogFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
	return *log
}
